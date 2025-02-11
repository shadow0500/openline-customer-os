import * as React from 'react';
import { SVGProps } from 'react';
const SvgBlockquote = (props: SVGProps<SVGSVGElement>) => (
  <svg
    className='blockquote_svg__svg-icon'
    style={{
      width: '1.0009765625em',
      height: '1em',
      verticalAlign: 'middle',
      fill: 'currentColor',
      overflow: 'hidden',
    }}
    viewBox='0 0 1025 1024'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path d='M224.992 448c123.712 0 224 100.288 224 224s-100.288 224-224 224-224-100.288-224-224L0 640c0-247.424 200.576-448 448-448v128c-85.472 0-165.824 33.28-226.272 93.728-11.648 11.648-22.24 24.032-31.84 37.024A226.597 226.597 0 0 1 224.992 448zm576 0c123.712 0 224 100.288 224 224s-100.288 224-224 224-224-100.288-224-224L576 640c0-247.424 200.576-448 448-448v128c-85.472 0-165.824 33.28-226.272 93.728-11.648 11.648-22.24 24.032-31.84 37.024A226.91 226.91 0 0 1 800.992 448z' />
  </svg>
);
export default SvgBlockquote;
