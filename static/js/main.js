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

    toJson() {
        return {
            senderID: this.senderID,
            event:    this.event,
            content:  this.content,
        }
    }

}

class WebSocketContainer {

    self;

    id = "";

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
        console.log(`Received Message:`, msg.data)
    }

    // Send a message
    sendMessage = (event, msg) => {

        if ( typeof event != 'string' ) {
            throw new Error(`Invalid parameter, please only provide strings events`);
        }

        // Build the send contents
        const sendContents = new WebSocketPacket(this.id, event, msg);

        // Send to server
        return this.self.send(JSON.stringify(sendContents));
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