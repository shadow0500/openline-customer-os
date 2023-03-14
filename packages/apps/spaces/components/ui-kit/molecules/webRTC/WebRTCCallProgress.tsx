import * as React from 'react';
import { useContext, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faMicrophone,
  faMicrophoneSlash,
  faPause,
  faPhoneSlash,
  faPlay,
  faRightLeft,
  faXmarkSquare,
} from '@fortawesome/free-solid-svg-icons';

import { WebRTCContext } from '../../../../context/web-rtc';
import styles from './web-rtc.module.scss';
import { Button, IconButton } from '../../atoms';
import { useRecoilValue } from 'recoil';
import { callParticipant } from '../../../../state';
import { Dialog } from 'primereact/dialog';
import { InputTextarea } from 'primereact/inputtextarea';
import { SuggestionList } from './SuggestionList';
import { useContactSuggestions } from '../../../../hooks/useContactSuggestions';
export const WebRTCCallProgress: React.FC<any> = () => {
  const [showRefer, setShowRefer] = useState(false);
  const [transferDest, setTransferDest] = useState('');
  const [referProgressString, setReferProgressString] = useState('');
  const [inRefer, setInRefer] = useState(false);
  const contactSuggestions = useContactSuggestions({ value: transferDest });
  const {
    inCall,
    isCallMuted,
    muteCall,
    unMuteCall,
    isCallOnHold,
    holdCall,
    unHoldCall,
    sendDtmf,
    hangupCall,
    ringing,
    transferCall,
  } = useContext(WebRTCContext) as any;
  const { identity } = useRecoilValue(callParticipant);
  const toggleMute = () => {
    if (isCallMuted) {
      unMuteCall();
    } else {
      muteCall();
    }
  };

  const toggleHold = () => {
    if (isCallOnHold) {
      unHoldCall();
    } else {
      holdCall();
    }
  };

  const getRows = () => {
    const makeButton = (number: string) => {
      return (
        <Button
          mode='secondary'
          key={'dtmf-' + number}
          onClick={() => {
            sendDtmf(number);
          }}
        >
          {number}
        </Button>
      );
    };

    const dialpad_matrix = new Array(4);
    for (let i = 0, digit = 1; i < 3; i++) {
      dialpad_matrix[i] = new Array(3);
      for (let j = 0; j < 3; j++, digit++) {
        dialpad_matrix[i][j] = makeButton(digit.toString());
      }
    }
    dialpad_matrix[3] = new Array(3);
    dialpad_matrix[3][0] = makeButton('*');
    dialpad_matrix[3][1] = makeButton('0');
    dialpad_matrix[3][2] = makeButton('#');
    const dialpad_rows = [];

    for (let i = 0; i < 4; i++) {
      dialpad_rows.push(
        <div key={'dtmf-row-' + i} className={styles.dialNumbersRow}>
          {dialpad_matrix[i]}
        </div>,
      );
    }

    return dialpad_rows;
  };

  if (!inCall) {
    return null;
  }
  const transfer = () => {
    setReferProgressString('');
    setInRefer(true);
    transferCall(transferDest);
  };
  return (
    <Dialog
      visible={inCall && !ringing}
      modal={false}
      className={styles.overlayContentWrapper}
      closable={false}
      closeOnEscape={false}
      draggable={false}
      onHide={() => console.log()}
    >
      <article>
        <h1 className={styles.sectionTitle}>In call with {identity}</h1>

        <div className={styles.dialNumbers}>{getRows()}</div>

        <div className={styles.actionButtonsRow}>
          <IconButton
            size='xxs'
            mode='primary'
            onClick={() => toggleMute()}
            icon={
              <FontAwesomeIcon
                icon={isCallMuted ? faMicrophone : faMicrophoneSlash}
              />
            }
          />

          <IconButton
            size='xxs'
            mode='primary'
            onClick={() => toggleHold()}
            icon={<FontAwesomeIcon icon={isCallOnHold ? faPlay : faPause} />}
          />

          <IconButton
            size='xxs'
            onClick={() => hangupCall()}
            mode='danger'
            icon={<FontAwesomeIcon icon={faPhoneSlash} />}
          />

          <IconButton
            size='xxs'
            onClick={() => setShowRefer(!showRefer)}
            mode='danger'
            icon={
              <FontAwesomeIcon icon={showRefer ? faXmarkSquare : faRightLeft} />
            }
          />
        </div>
      </article>
      {showRefer && (
        <>
          <div>
            <div className='w-full text-center align-items-center mb-3'>
              <InputTextarea
                className='mr-2'
                value={transferDest}
                onChange={(e) => setTransferDest(e.target.value)}
                autoResize
                rows={1}
                placeholder='Transfer to'
                onKeyPress={(e) => {
                  if (e.shiftKey && e.key === 'Enter') {
                    return true;
                  }
                  if (e.key === 'Enter') {
                    e.preventDefault();
                  }
                }}
                style={{
                  borderColor: 'black', //Do not set as none!! It breaks InputTextarea autoResize
                  boxShadow: 'none',
                }}
              />
              <span className='h-full align-items-top'>
                <Button
                  onClick={transfer}
                  className='p-button-success h-full mr-2'
                >
                  <FontAwesomeIcon icon={faRightLeft} className='mr-2' />
                </Button>
              </span>
            </div>
            <div>
              <SuggestionList
                currentValue={transferDest}
                getSuggestions={contactSuggestions}
                setCurrentValue={setTransferDest}
              ></SuggestionList>
            </div>
          </div>
        </>
      )}
      {inRefer && (
        <div
          style={{
            position: 'absolute',
            zIndex: 2000,
            width: '100%',
            height: '100%',
            top: '0%',
            background: '#FFFFFFFF',
          }}
        >
          <div
            style={{
              margin: 0,
              position: 'absolute',
              top: '50%',
              transform: 'translateY(-50%)',
              width: '100%',
            }}
          >
            <div className='w-full text-center align-items-center mb-3'>
              Transfering call to: {transferDest}
            </div>
            <div
              key='referProgress'
              className='w-full text-center align-items-center mb-3'
            >
              {referProgressString}
            </div>
            <div className='w-full text-center align-items-center mb-3'>
              <FontAwesomeIcon icon={faRightLeft} className='mr-2' />
            </div>
          </div>
        </div>
      )}
    </Dialog>
  );
};
