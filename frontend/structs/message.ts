interface iMessage {
    id: number;
    text: string;
    sender: string;
    timestamp: number | string;
    user: string | null
}

interface Send {
    user: string,
    text: string,
    // timestamp: 
}