class WebSocketPacket {

    senderID = null;
    event    = null;
    content  = null;

    /**
     * 
     * @param {String} senderID The id of the sender
     * @param {String} event The action that we are sending
     * @param {*} content The content we are sending
     */
    constructor(senderID, event, content) {
        this.senderID = senderID
        this.event    = event
        this.content  = content
    }

    toJson = () => {
        return {
            senderID: this.senderID,
            event:    this.event,
            content:  this.content,
        }
    }

}

class WebSocketContainer {

    self = null;
    id   = null;

    constructor(webSocketURL) {
        // set senderID
        this.id = generateGuid();
        // Spawn self
        this.self = new WebSocket(webSocketURL);
        this.self.onclose = this.onClose;
        this.self.onopen  = this.onConnected;
        this.self.onmessage = this.onMessage;
    }

    onConnected = () => {
        console.log(`Web Socket connection established`);
        this.sendMessage(`connected`, null);
        // Send self information to the database
    }

    onClose = () => {
        console.warn(`Web Socket connection closed`);
    }

    onMessage = (msg) => {

        // Parse the message
        let parsed = JSON.parse(msg.data);

        // Check what kind of event this is
        if ( parsed['event'] == 'update_database' ) {
            this.onDatabaseUpdate(parsed['content']);
        }

    }

    // Send a message
    sendMessage = (event, msg) => {

        if ( typeof event != 'string' ) {
            throw new Error(`Invalid parameter, please only provide strings events`);
        }

        // Build the send contents
        const sendContents = new WebSocketPacket(this.id, event, msg).toJson();

        // Send to server
        return this.self.send(JSON.stringify(sendContents));
    }

    onDatabaseUpdate = data => {

        // Go through all channels
        for ( let [channelName, d] of Object.entries(data) ) {

            // Check if the channel exists
            let myChannel = document.querySelector(`*[channel-id="${channelName}"]`);
            if ( !myChannel ) {
                myChannel = this.createChannel(channelName, d);
            }

            this.updateChannel(myChannel, d)
        }

    }

    /**
     * 
     * @param {HTMLElement} myChannel 
     * @param {Array} cdata 
     */
    updateChannel = (myChannel, cdata) => {
        // Update the message indicator
        myChannel.querySelector(`.gossip-indicator`).textContent = cdata.length;

        this.updateChannelMessages(myChannel, cdata);
    }

    updateChannelMessages = (myChannel, cdata) => {
        // Check if my channel is actually selected
        if ( myChannel.classList.contains(`selected`) ) {
            // Get the main container
            const msgContainer = document.getElementById(`gossip-messages`);
            // Clean old logs
            msgContainer.querySelectorAll(`div`).forEach( e => {
                e.remove();
            })

            // Create the new messages
            for ( let i of cdata ) {
                const row = document.createElement(`div`);
                // Append the message
                msgContainer.prepend(row);

                row.classList.add(`gossip-message`);
                row.innerHTML = `<span class="ts">${i['timestamp']}</span><span class="msg">${i['content']}</span>`;
            }

        }
    }

    createChannel = (cname, cdata) => {

        const templateMessage = document.getElementById(`template-message`);

        if ( !!!templateMessage ) {
            throw new Error(`Could not find the template message`);
        }

        /**
         * @type {HTMLElement}
         */
        const myChannel = templateMessage.cloneNode(true);
        // Remove ID
        myChannel.removeAttribute(`id`);
        // Change name
        myChannel.querySelector(`.channel-name`).textContent = cname;
        // Set the message indicator
        myChannel.querySelector(`.gossip-indicator`).textContent = cdata.length;
        // Set the id of the channel
        myChannel.setAttribute(`channel-id`, cname);

        myChannel.addEventListener(`click`, e => {
            document.querySelectorAll(`.gossip-room-select.selected`).forEach( e => {
                e.classList.remove(`selected`);
            })

            myChannel.classList.add(`selected`);

            this.updateChannelMessages(myChannel, cdata);
        })


        // Append the element
        templateMessage.parentElement.appendChild(myChannel);

        return myChannel;
    }

}






/**
 * Generate a GUID
 */
 function generateGuid() {
    let result = '';
    for (let j = 0; j < 32; j++) {
        if (j == 8 || j == 12 || j == 16 || j == 20)
            result = result + '-';
        i = Math.floor(Math.random() * 16).toString(16).toUpperCase();
        result = result + i;
    }
    return result;
}