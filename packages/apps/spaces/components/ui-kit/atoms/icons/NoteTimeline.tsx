import * as React from 'react';
import { SVGProps } from 'react';
const SvgNoteTimeline = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 144 58'
    fill='none'
    {...props}
  >
    <g filter='url(#note-timeline_svg__a)'>
      <rect
        x={4}
        width={136}
        height={50}
        rx={4}
        fill='#fff'
        shapeRendering='crispEdges'
      />
      <path
        d='M4.5 6V4A3.5 3.5 0 0 1 8 .5h128a3.5 3.5 0 0 1 3.5 3.5v2H4.5z'
        fill='#f4e69c'
      />
      <path
        d='M89.665 26.602V33h-.853l-3.221-4.935V33h-.848v-6.398h.848l3.234 4.948v-4.948h.839zm1.165 4.074v-.101c0-.343.05-.661.149-.954s.243-.552.431-.769a1.92 1.92 0 0 1 .681-.51c.267-.123.566-.185.897-.185s.634.062.901.185a1.91 1.91 0 0 1 .686.51c.191.217.335.473.435.769s.149.611.149.954v.101c0 .343-.05.661-.149.954a2.28 2.28 0 0 1-.435.769 1.991 1.991 0 0 1-.681.51 2.14 2.14 0 0 1-.897.18 2.17 2.17 0 0 1-.901-.18 2.04 2.04 0 0 1-.686-.51 2.32 2.32 0 0 1-.431-.769 2.936 2.936 0 0 1-.149-.954zm.813-.101v.101c0 .237.028.461.084.672s.139.393.251.554a1.27 1.27 0 0 0 .426.382 1.24 1.24 0 0 0 .593.136c.223 0 .418-.045.585-.136a1.23 1.23 0 0 0 .422-.382 1.77 1.77 0 0 0 .25-.554c.059-.211.088-.435.088-.672v-.101c0-.234-.029-.456-.088-.664a1.71 1.71 0 0 0-.255-.558c-.111-.164-.252-.293-.422-.387s-.363-.141-.589-.141a1.2 1.2 0 0 0-.589.141c-.167.094-.308.223-.422.387s-.195.347-.251.558-.084.429-.084.664zm6.526-2.329v.624h-2.571v-.624h2.571zm-1.701-1.156h.813v4.733c0 .161.025.283.075.365s.114.136.193.163a.8.8 0 0 0 .255.04 1.34 1.34 0 0 0 .211-.018l.171-.035.004.664a1.89 1.89 0 0 1-.255.057 1.94 1.94 0 0 1-.374.031c-.199 0-.382-.04-.549-.119s-.3-.211-.4-.395-.145-.439-.145-.756v-4.728zm4.628 5.998c-.331 0-.632-.056-.901-.167a2.04 2.04 0 0 1-.69-.479 2.13 2.13 0 0 1-.439-.729 2.664 2.664 0 0 1-.154-.923v-.185c0-.387.057-.731.171-1.033s.27-.562.466-.773.419-.371.668-.479a1.92 1.92 0 0 1 .773-.163c.34 0 .633.059.879.176a1.6 1.6 0 0 1 .611.492c.158.208.275.454.352.738a3.53 3.53 0 0 1 .114.923v.365h-3.551v-.664h2.738v-.061a2.04 2.04 0 0 0-.132-.615 1.1 1.1 0 0 0-.352-.492c-.161-.129-.38-.193-.659-.193a1.11 1.11 0 0 0-.51.119 1.085 1.085 0 0 0-.399.343 1.72 1.72 0 0 0-.26.558 2.83 2.83 0 0 0-.092.76v.185a2.15 2.15 0 0 0 .092.637c.064.196.157.369.277.518s.271.267.444.352.375.127.598.127c.287 0 .53-.059.729-.176s.374-.274.523-.47l.492.391a2.26 2.26 0 0 1-.391.444c-.158.141-.353.255-.584.343s-.5.132-.813.132z'
        fill='#1d1d1d'
      />
      <rect
        x={4.25}
        y={0.25}
        width={135.5}
        height={49.5}
        rx={3.75}
        stroke='#9b9b9b'
        strokeWidth={0.5}
        shapeRendering='crispEdges'
      />
    </g>
    <defs>
      <filter
        id='note-timeline_svg__a'
        x={0}
        y={0}
        width={144}
        height={58}
        filterUnits='userSpaceOnUse'
        colorInterpolationFilters='sRGB'
      >
        <feFlood floodOpacity={0} result='A' />
        <feColorMatrix
          in='SourceAlpha'
          values='0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0'
          result='B'
        />
        <feOffset dy={4} />
        <feGaussianBlur stdDeviation={2} />
        <feComposite in2='B' operator='out' />
        <feColorMatrix values='0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0' />
        <feBlend in2='A' />
        <feBlend in='SourceGraphic' />
      </filter>
    </defs>
  </svg>
);
export default SvgNoteTimeline;
