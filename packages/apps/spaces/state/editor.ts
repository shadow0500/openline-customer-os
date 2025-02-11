import { atom } from 'recoil';

export enum EditorMode {
  Note = 'NOTE',
  Email = 'EMAIL',
  Chat = 'CHAT',
  PhoneCall = 'PHONE_CALL',
}
export interface EmailMode {
  handleSubmit?: (
    data: any,
    onSuccess: () => void,
    destination: Array<string>,
    respondTo: null | string,
  ) => Promise<any>;
  subject: string;
  to: Array<string>;
  respondTo: null | string;
}

export const editorMode = atom({
  key: 'editor', // unique ID (with respect to other atoms/selectors)
  default: {
    mode: EditorMode.Note,
  }, // default value (aka initial value)
});
export const editorEmail = atom<EmailMode>({
  key: 'editorEmail', // unique ID (with respect to other atoms/selectors)
  default: {
    to: [],
    subject: '',
    respondTo: null,
  },
});

export const showLegacyEditor = atom({
  key: 'showLegacyEditor', // unique ID (with respect to other atoms/selectors)
  default: false,
});
