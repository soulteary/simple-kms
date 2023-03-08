package home

var Template = []byte(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <!-- css design by @LukyVj -->
    <!-- https://codepen.io/LukyVj/pen/YzOXepM -->
    <style>
      @property --hue {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --rotate {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --bg-y {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --bg-x {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --glow-translate-y {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --bg-size {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --glow-opacity {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --glow-blur {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      @property --glow-scale {
        syntax: "<number>";
        inherits: true;
        initial-value: 2;
      }
      @property --glow-radius {
        syntax: "<number>";
        inherits: true;
        initial-value: 2;
      }
      @property --white-shadow {
        syntax: "<number>";
        inherits: true;
        initial-value: 0;
      }
      :root {
        --debug: 0;
        --supported: 0;
        --not-supported: 0;
        --card-color: hsl(260deg 100% 3%);
        --text-color: hsl(260deg 10% 55%);
        --card-radius: 3.6vw;
        --card-width: 35vw;
        --border-width: 3px;
        --bg-size: 1;
        --hue: 0;
        --hue-speed: 1;
        --rotate: 0;
        --animation-speed: 4s;
        --interaction-speed: 0.55s;
        --glow-scale: 1.5;
        --scale-factor: 1;
        --glow-blur: 6;
        --glow-opacity: 1;
        --glow-radius: 100;
        --glow-rotate-unit: 1deg;
      }

      html,
      body {
        height: 100%;
        width: 100%;
        padding: 0;
        margin: 0;
      }

      *,
      *:before,
      *:after {
        outline: calc(var(--debug) * 1px) red dashed;
      }

      body {
        background-color: var(--card-color);
        display: flex;
        align-items: center;
        justify-content: center;
        font-family: sans-serif;
      }

      body > div {
        width: var(--card-width);
        width: min(480px, var(--card-width));
        aspect-ratio: 1.5/1;
        color: white;
        margin: auto;
        display: flex;
        align-items: center;
        justify-content: center;
        position: relative;
        z-index: 2;
        border-radius: var(--card-radius);
        cursor: pointer;
      }
      body > div:hover > div {
        mix-blend-mode: darken;
        --text-color: white;
        box-shadow: 0 0 calc(var(--white-shadow) * 1vw) calc(var(--white-shadow) * 0.15vw) rgba(255, 255, 255, 0.2);
        animation: shadow-pulse calc(var(--animation-speed) * 2) linear infinite;
      }
      body > div:hover > div:before {
        --bg-size: 15;
        animation-play-state: paused;
        transition: --bg-size var(--interaction-speed) ease;
      }
      body > div:hover .glow {
        --glow-blur: 1.5;
        --glow-opacity: 0.6;
        --glow-scale: 2.5;
        --glow-radius: 0;
        --rotate: 900;
        --glow-rotate-unit: 0;
        --scale-factor: 1.25;
        animation-play-state: paused;
      }
      body > div:hover .glow:after {
        --glow-translate-y: 0;
        animation-play-state: paused;
        transition: --glow-translate-y 0s ease, --glow-blur 0.05s ease, --glow-opacity 0.05s ease, --glow-scale 0.05s ease, --glow-radius 0.05s ease;
      }
      body > div:before,
      body > div:after {
        content: "";
        display: block;
        position: absolute;
        width: 100%;
        height: 100%;
        border-radius: var(--card-radius);
      }
      body > div > div {
        position: absolute;
        width: 100%;
        height: 100%;
        background: var(--card-color);
        border-radius: calc(calc(var(--card-radius) * 0.9));
        display: flex;
        align-items: center;
        justify-content: center;
        font-weight: 800;
        text-transform: uppercase;
        font-stretch: 150%;
        font-size: 18px;
        font-size: clamp(1.5vw, 1.5vmin, 32px);
        color: var(--text-color);
        padding: calc(var(--card-width) / 8);
      }
      body > div > div span {
        display: inline-block;
        padding: 0.25em;
        border-radius: 4px;
        background: var(--text-color);
        color: black;
        margin-right: 8px;
        font-weight: 900;
      }
      body > div > div:before {
        content: "";
        display: block;
        position: absolute;
        width: 100%;
        height: 100%;
        border-radius: calc(calc(var(--card-radius) * 0.9));
        box-shadow: 0 0 20px black;
        mix-blend-mode: color-burn;
        z-index: -1;
        background: #292929
          radial-gradient(
            30% 30% at calc(var(--bg-x) * 1%) calc(var(--bg-y) * 1%),
            hsl(calc(calc(var(--hue) * var(--hue-speed)) * 1deg), 100%, 90%) calc(0% * var(--bg-size)),
            hsl(calc(calc(var(--hue) * var(--hue-speed)) * 1deg), 100%, 80%) calc(20% * var(--bg-size)),
            hsl(calc(calc(var(--hue) * var(--hue-speed)) * 1deg), 100%, 60%) calc(40% * var(--bg-size)),
            transparent 100%
          );
        width: calc(100% + var(--border-width));
        height: calc(100% + var(--border-width));
        animation: hue-animation var(--animation-speed) linear infinite, rotate-bg var(--animation-speed) linear infinite;
        transition: --bg-size var(--interaction-speed) ease;
      }
      body > div .glow {
        --glow-translate-y: 0;
        display: block;
        position: absolute;
        width: calc(var(--card-width) / 5);
        height: calc(var(--card-width) / 5);
        animation: rotate var(--animation-speed) linear infinite;
        transform: rotateZ(calc(var(--rotate) * var(--glow-rotate-unit)));
        transform-origin: center;
        border-radius: calc(var(--glow-radius) * 10vw);
      }
      body > div .glow:after {
        content: "";
        display: block;
        z-index: -2;
        filter: blur(calc(var(--glow-blur) * 10px));
        width: 130%;
        height: 130%;
        left: -15%;
        top: -15%;
        background: hsl(calc(calc(var(--hue) * var(--hue-speed)) * 1deg), 100%, 60%);
        position: relative;
        border-radius: calc(var(--glow-radius) * 10vw);
        animation: hue-animation var(--animation-speed) linear infinite;
        transform: scaleY(calc(var(--glow-scale) * var(--scale-factor) / 1.1)) scaleX(calc(var(--glow-scale) * var(--scale-factor) * 1.2)) translateY(calc(var(--glow-translate-y) * 1%));
        opacity: var(--glow-opacity);
      }

      @keyframes shadow-pulse {
        0%,
        24%,
        46%,
        73%,
        96% {
          --white-shadow: 0.5;
        }
        12%,
        28%,
        41%,
        63%,
        75%,
        82%,
        98% {
          --white-shadow: 2.5;
        }
        6%,
        32%,
        57% {
          --white-shadow: 1.3;
        }
        18%,
        52%,
        88% {
          --white-shadow: 3.5;
        }
      }
      @keyframes rotate-bg {
        0% {
          --bg-x: 0;
          --bg-y: 0;
        }
        25% {
          --bg-x: 100;
          --bg-y: 0;
        }
        50% {
          --bg-x: 100;
          --bg-y: 100;
        }
        75% {
          --bg-x: 0;
          --bg-y: 100;
        }
        100% {
          --bg-x: 0;
          --bg-y: 0;
        }
      }
      @keyframes rotate {
        from {
          --rotate: -70;
          --glow-translate-y: -65;
        }
        25% {
          --glow-translate-y: -65;
        }
        50% {
          --glow-translate-y: -65;
        }
        60%,
        75% {
          --glow-translate-y: -65;
        }
        85% {
          --glow-translate-y: -65;
        }
        to {
          --rotate: calc(360 - 70);
          --glow-translate-y: -65;
        }
      }
      @keyframes hue-animation {
        0% {
          --hue: 0;
        }
        100% {
          --hue: 360;
        }
      }
    </style>
  </head>
  <body>
    <div role="button">
      <span class="glow"></span>
      <div><span>Simple</span>KMS authentication</div>
    </div>
  </body>
</html>
`)
