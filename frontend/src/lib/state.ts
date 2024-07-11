import { writable } from 'svelte/store';
import { type Room, type ChatMessage } from './interfaces'

export const room = writable({} as Room);
export const chatMessages = writable([] as ChatMessage[]);
