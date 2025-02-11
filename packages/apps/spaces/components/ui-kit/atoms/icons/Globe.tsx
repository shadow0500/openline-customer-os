import * as React from 'react';
import { SVGProps } from 'react';
const SvgGlobe = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M12 3a9 9 0 0 0-6.364 15.364A9 9 0 0 0 21 12a9 9 0 0 0-9-9zm7.46 8.25H16.7a13 13 0 0 0-2.94-6.53 7.52 7.52 0 0 1 5.7 6.53zm-10.65 1.5h6.38A13.18 13.18 0 0 1 12 19.1a13.18 13.18 0 0 1-3.19-6.35zm0-1.5A13.18 13.18 0 0 1 12 4.9a13.18 13.18 0 0 1 3.19 6.35H8.81zm1.43-6.53a13 13 0 0 0-2.94 6.53H4.54a7.52 7.52 0 0 1 5.7-6.53zm-5.7 8H7.3a13 13 0 0 0 2.94 6.53 7.52 7.52 0 0 1-5.7-6.5v-.03zm9.22 6.53a13 13 0 0 0 2.94-6.53h2.76a7.52 7.52 0 0 1-5.7 6.56v-.03z'
      fill='currentColor'
    />
  </svg>
);
export default SvgGlobe;
