import * as React from 'react';
import { SVGProps } from 'react';
const SvgBox = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 16 17'
    fill='none'
    {...props}
  >
    <path
      fill='currentColor'
      d='M1.5 13a1 1 0 0 1-.6-.5l-.2-.9V5.4c0-.3 0-.5.2-.7l.4-.4L7 1.1A2 2 0 0 1 8 1c.3 0 .7 0 1 .2l5.7 3.2.4.4c.2.2.2.4.2.7v6.2c0 .4 0 .7-.2.9a1 1 0 0 1-.6.5l-6 3.3-.5.2a1 1 0 0 1-.5-.2l-6-3.3zm.5-.5 5.6 3V9L1.4 5.6v6l.1.5.5.4zm12 0 .5-.4V5.6L8.5 9v6.6l5.6-3.1zM8 8.3l2.8-1.5-6.3-3.5L1.8 5 8 8.3zm3.6-2L14.2 5 8.6 1.7c-.4-.2-.8-.2-1.2 0L5.3 3l6.3 3.5z'
    />
  </svg>
);
export default SvgBox;
