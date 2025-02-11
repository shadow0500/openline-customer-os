import * as React from 'react';
import { SVGProps } from 'react';
const SvgList = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    xmlnsXlink='http://www.w3.org/1999/xlink'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <g fill='currentColor'>
      <use xlinkHref='#list_svg__a' />
      <use xlinkHref='#list_svg__a' y={-4.5} />
      <use xlinkHref='#list_svg__a' y={4.5} />
      <path d='M5 8.5a1 1 0 0 1-.38-.07 1.46 1.46 0 0 1-.33-.22A1 1 0 0 1 4 7.5a1.05 1.05 0 0 1 .29-.71.93.93 0 0 1 .33-.21 1 1 0 0 1 .76 0 1 1 0 0 1 .33.21A1.05 1.05 0 0 1 6 7.5a1 1 0 0 1-.29.71 1.46 1.46 0 0 1-.33.22A1 1 0 0 1 5 8.5zM5 13a1 1 0 0 1-.38-.08 1.15 1.15 0 0 1-.33-.21A1 1 0 0 1 4 12a1.05 1.05 0 0 1 .29-.71 1.15 1.15 0 0 1 .33-.21A1 1 0 0 1 5.2 11l.18.06.18.09c.052.037.102.077.15.12A1.05 1.05 0 0 1 6 12a1 1 0 0 1-1 1zm0 4.5a1 1 0 0 1-.38-.07 1.46 1.46 0 0 1-.33-.22 1.15 1.15 0 0 1-.21-.33.94.94 0 0 1 0-.76 1.15 1.15 0 0 1 .21-.33 1 1 0 0 1 1.09-.21 1 1 0 0 1 .33.21 1.15 1.15 0 0 1 .21.33.94.94 0 0 1 0 .76 1.15 1.15 0 0 1-.21.33 1 1 0 0 1-.71.29z' />
    </g>
    <defs>
      <path
        id='list_svg__a'
        d='M19 12.75H8a.75.75 0 0 1-.75-.75.75.75 0 0 1 .75-.75h11a.75.75 0 0 1 .75.75.75.75 0 0 1-.75.75z'
      />
    </defs>
  </svg>
);
export default SvgList;
