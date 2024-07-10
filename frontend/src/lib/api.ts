import { type Room } from "../lib/interfaces"


export async function createRoom() {
    try {
        let url = `${import.meta.env.VITE_API_URL}/rooms/new`;
        const response = await fetch(url, {method: "POST"});
        if (!response.ok) {
            throw new Error("Failed to fetch data")
        }
        const data = await response.json();
        console.log(data)


        const room: Room = data as Room;
        return room;
    } catch (error) {
        console.log("Error fetching data: ", error)
        throw error;
    }
}
