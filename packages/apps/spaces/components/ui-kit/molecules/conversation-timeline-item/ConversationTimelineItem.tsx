import React, { useEffect, useRef, useState } from 'react';
import VoiceWave from '@spaces/atoms/icons/VoiceWave';
import ChevronUp from '@spaces/atoms/icons/ChevronUp';
import MessageIcon from '@spaces/atoms/icons/MessageIcon';
import ChevronDown from '@spaces/atoms/icons/ChevronDown';
import Phone from '@spaces/atoms/icons/Phone';
import ArrowRight from '@spaces/atoms/icons/ArrowRight';
import ArrowLeft from '@spaces/atoms/icons/ArrowLeft';
import { Tooltip } from '@spaces/atoms/tooltip';
import { TimelineItem } from '@spaces/atoms/timeline-item';
import { AnalysisContent } from '@spaces/atoms/message/AnalysisContent';
import { TranscriptContent } from '@spaces/atoms/message/TranscriptContent';

import classNames from 'classnames';
import {
  ConversationPartyEmail,
  ConversationPartyPhone,
} from './ConversationParty';
import styles from './conversation-timeline-item.module.scss';
import { DataSource } from '../../../../graphQL/__generated__/generated';

interface Content {
  dialog: {
    type?: string;
    mimetype: string;
    body: string;
  };
}

interface TranscriptElement {
  party: any;
  text: string;
  file_id?: string;
}

interface TranscriptV2 {
  transcript: Array<TranscriptElement>;
  file_id?: string;
}

interface Props {
  createdAt: string;
  content: Content | undefined;
  transcript: Array<TranscriptElement>;
  type: string;
  mode: 'PHONE_CALL' | 'CHAT';
  id: string;
  contentType?: string;
  source: DataSource;
}

interface DataStateI {
  firstSendIndex: null | number;
  firstReceivedIndex: null | number;
  initiator: 'left' | 'right';
}

const getTranscript = (
  transcript: TranscriptV2 | Array<any>,
  contentType: string | undefined,
): Array<any> => {
  if (contentType === 'application/x-openline-transcript') {
    return transcript as Array<any>;
  } else if (contentType === 'application/x-openline-transcript-v2') {
    const transcript2 = transcript as TranscriptV2;
    return transcript2.transcript;
  }
  return transcript as Array<any>;
};

const getFileId = (
  transcript: TranscriptV2 | Array<any>,
  contentType: string | undefined,
): string | undefined => {
  if (contentType === 'application/x-openline-transcript') {
    return undefined;
  } else if (contentType === 'application/x-openline-transcript-v2') {
    const transcript2 = transcript as TranscriptV2;
    return transcript2.file_id;
  }
  return undefined;
};

export const ConversationTimelineItem: React.FC<Props> = ({
  createdAt,
  content,
  transcript = [],
  type,
  mode = 'PHONE_CALL',
  id,
  contentType,
  source,
}) => {
  const messagesContainerRef = useRef<HTMLDivElement>(null);
  const summaryRef = useRef<HTMLDivElement>(null);

  const [data, setData] = useState<DataStateI>({
    firstSendIndex: null,
    firstReceivedIndex: null,
    initiator: 'left',
  });
  const [summaryExpanded, setSummaryExpanded] = useState(false);
  const handleToggleExpanded = () => {
    setSummaryExpanded(!summaryExpanded);
    if (summaryRef?.current && summaryExpanded) {
      summaryRef?.current?.scrollIntoView({ behavior: 'smooth' });
    }
  };

  useEffect(() => {
    if (data.firstSendIndex === null) {
      const left = getTranscript(transcript, contentType).findIndex(
        (e: TranscriptElement) => e?.party?.tel,
      );
      const right = getTranscript(transcript, contentType).findIndex(
        (e: TranscriptElement) => e?.party?.mailto,
      );

      setData({
        firstSendIndex: left,
        firstReceivedIndex: right,
        initiator: left === 0 ? 'left' : 'right',
      });
    }
  }, []);

  // fixme for some reason it does not work whe put in state
  const left = getTranscript(transcript, contentType)?.find(
    (e: TranscriptElement) => e?.party?.tel,
  );
  const right = getTranscript(transcript, contentType)?.find(
    (e: TranscriptElement) => e?.party?.mailto,
  );

  //const right=false, left = false;
  return (
    <div className='flex flex-column w-full'>
      <TimelineItem source={source} first createdAt={createdAt}>
        <div
          className={classNames(styles.contentWrapper, {
            [styles.expanded]: summaryExpanded,
          })}
        >
          {type === 'summary' && (
            <>
              <div className='flex flex-column w-full'>
                <div className={styles.summary} ref={summaryRef}>
                  <div
                    className={classNames(styles.left, {
                      [styles.initiator]: data.initiator === 'left',
                    })}
                  >
                    <div className={styles.callPartyData}>
                      <ConversationPartyPhone tel={left?.party.tel} />

                      <div className={styles.iconsWrapper}>
                        {getTranscript(transcript, contentType)?.[0]?.party
                          .tel && (
                          <>
                            <VoiceWave height={20} />
                            <ArrowRight height={20} />
                          </>
                        )}
                      </div>
                    </div>
                  </div>

                  <div
                    className={classNames(styles.right, {
                      [styles.initiator]: data.initiator === 'right',
                    })}
                  >
                    <div className={styles.callPartyData}>
                      <div className={styles.iconsWrapper}>
                        {!getTranscript(transcript, contentType)?.[0]?.party
                          .tel && (
                          <>
                            <ArrowLeft height={20} />
                            <VoiceWave height={20} />
                          </>
                        )}
                      </div>
                      <ConversationPartyEmail
                        email={(right?.party.mailto || '').toLowerCase()}
                      />
                    </div>
                  </div>
                </div>
              </div>

              {content && (
                <Tooltip
                  content={content.dialog?.body || ''}
                  target={`#phone-summary-${id}`}
                  position='bottom'
                  showDelay={300}
                  autoHide={false}
                />
              )}

              <button
                id={`phone-summary-${id}`}
                className={styles.folderTab}
                role='button'
                onClick={handleToggleExpanded}
              >
                {summaryExpanded ? (
                  <ChevronUp height={16} color='#3A8745' />
                ) : (
                  <ChevronDown height={16} color='#3A8745' />
                )}

                {content && (
                  <span>
                    Summary: <AnalysisContent analysis={content.dialog} />
                  </span>
                )}
              </button>
            </>
          )}

          <section
            ref={messagesContainerRef}
            className={classNames(styles.transcriptionContainer, {
              [styles.transcriptionContainerOpen]: summaryExpanded,
            })}
            style={{
              maxHeight: summaryExpanded
                ? `${messagesContainerRef?.current?.scrollHeight}px`
                : 0,
            }}
          >
            <div className={styles.messages}>
              <TranscriptContent
                contentType={contentType}
                messages={getTranscript(transcript, contentType)}
                firstIndex={{
                  received: 0,
                  send: 1,
                }}
              >
                {mode === 'CHAT' ? (
                  <MessageIcon height={24} />
                ) : (
                  <Phone height={24} />
                )}
              </TranscriptContent>
              {getFileId(transcript, contentType) && (
                <video controls style={{ width: '100%' }}>
                  <source
                    src={`/fs/file/${getFileId(
                      transcript,
                      contentType,
                    )}/download?inline=true`}
                  />
                </video>
              )}
            </div>
          </section>
        </div>
      </TimelineItem>
    </div>
  );
};
