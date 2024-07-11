import { version as uuidVersion } from 'uuid';
import { validate as uuidValidate } from 'uuid';

export function uuidValidateV4(uuid: string) {
  return uuidValidate(uuid) && uuidVersion(uuid) === 4;
}

export function getCookieValue(name: string) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop()?.split(';').shift();
    return null;
}
