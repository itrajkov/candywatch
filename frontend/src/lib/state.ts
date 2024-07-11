import { writable } from 'svelte/store';
import { type Room } from './interfaces'

export const room = writable({} as Room);
