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
        // Spawn self
        this.self = new WebSocket(webSocketURL);
        this.self.onclose = this.onClose;
        this.self.onopen  = this.onConnected;
        this.self.onmessage = this.onMessage;
    }

    onConnected = () => {
        console.log(`Web Socket connection established`);
        this.sendMessage(new WebSocketPacket(this.id, `connected`, {
            test: [ 1, 2, 3 ]
        }));
        // Send self information to the database
    }

    onClose = () => {
        console.warn(`Web Socket connection closed`);
    }

    onMessage = (msg) => {
        console.log(`Received Message:`, msg.data)
    }

    // Send a message
    sendMessage = (sendContents) => {

        if ( typeof sendContents == 'object' ) {
            if ( sendContents.constructor == WebSocketPacket ) {
                // Send to server
                return this.self.send(JSON.stringify(sendContents));
            }
        }

        throw new Error(`Invalid sendContents type, must be of type WebSocketPacket`);
    }

}