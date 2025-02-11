import * as React from 'react';
import styles from './message.module.scss';
import sanitizeHtml from 'sanitize-html';
import { TranscriptContent } from './TranscriptContent';
import linkifyHtml from 'linkify-html';
import { ReactNode } from 'react';

interface Content {
  type?: string;
  mimetype: string;
  body: string;
}

interface AnalysisContentProps {
  analysis: Content;
  children?: ReactNode;
}

export const AnalysisContent: React.FC<AnalysisContentProps> = ({
  analysis,
  children,
}) => {
  // Currently only used for Summary
  if (analysis?.mimetype === 'text/plain') {
    return <>{analysis.body}</>;
  }

  if (analysis?.mimetype === 'text/html') {
    return (
      <div
        className={`text-overflow-ellipsis ${styles.emailContent}`}
        dangerouslySetInnerHTML={{
          __html: sanitizeHtml(
            linkifyHtml(analysis.body, {
              defaultProtocol: 'https',
              rel: 'noopener noreferrer',
            }),
          ),
        }}
      ></div>
    );
  }

  return null;
};
