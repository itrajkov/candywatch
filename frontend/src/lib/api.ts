import { type Room } from "../lib/interfaces"
import { uuidValidateV4 } from "./util";

export async function createRoom(): Promise<Room> {
    try {
        let url = `${import.meta.env.VITE_API_URL}/rooms/new`;
        const response = await fetch(url, {
            method: "POST",
            credentials: 'include'
        });
        if (!response.ok) {
            throw new Error("Failed to create a room.")
        }
        const data = await response.json();
        const room: Room = data as Room;
        return room;
    } catch (error) {
        console.log("Failed to create a room:", error)
        throw error;
    }
}

export async function joinRoom(roomId: string): Promise<Room> {
    if (!uuidValidateV4(roomId)) {
        throw new Error("Not a valid UUID4.")
    }

    try {
        let url = `${import.meta.env.VITE_API_URL}/rooms/${roomId}/join`;
        const response = await fetch(url, {
            method: "POST", credentials: 'include'
        });
        if (!response.ok) {
            throw new Error("Failed to join room.")
        }
        const data = await response.json();
        const room: Room = data as Room;
        return room;
    } catch (error) {
        console.log("Failed to join room: ", error)
        throw error;
    }
}
