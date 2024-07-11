export interface Room {
    id: string;
    users: string[];
}

export interface ChatMessage {
    userId: string;
    content: string;
}
