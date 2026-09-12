<template>
<div class="tts-container" v-loading="loading">
    <label class="tts-label">
        {{$t('modelSquare.outputAduio')}}
    </label>
    <div class="component-wrapper" v-show="!loading">
        <div class="waveform-container"><div id="waveform" ref="waveform"></div></div>
        <div class="timestamps">
            <time id="time">{{currentTime}}</time>
            <div>
                <time id="duration">{{duration}}</time>
            </div>
        </div>
        <div class="waveform-controls">
            <div class="control-wrapper">
                <div class="volume-icon" @click="offVolume">
                    <i class="ri-volume-up-line" v-if="volumeValue && onVolume"></i>
                    <i class="ri-volume-mute-line" v-else></i>
                </div>
                <input class="volume-slider" v-model="volumeValue" @input="sliderChange($event)" type="range" min="0" max="1" step="0.01">
            </div>
            <div class="play-pause-icon">
                <div class="skip-left-icon" @click="skipBack"><i class="ri-rewind-fill 1"></i></div>
                <div class="play-icon" @click="playPause" v-if="playing"><i class="ri-play-fill"></i></div>
                <div class="pause-icon" @click="playPause" v-else><i class="ri-pause-line"></i></div>
                <div class="skip-right-icon" @click="skipForward"><i class="ri-speed-fill"></i></div>
            </div>
            <div class="download-icon" @click="download"><i class="ri-download-line"></i></div>
        </div>
    </div>
</div>
  
</template>

<script>
import WaveSurfer from "wavesurfer.js";
const options = {
  /** The height of the waveform in pixels */
  height: 80,
  /** The width of the waveform in pixels or any CSS value; defaults to 100% */
  //   width: 300,
  /** Render each audio channel as a separate waveform */
  splitChannels: false,
  /** Stretch the waveform to the full height */
  normalize: false,
  /** The color of the waveform */
  waveColor: '#9ca3af',
  /** The color of the progress mask */
  progressColor: '#f29849',
  /** The color of the playpack cursor */
  cursorColor: '#ddd5e9',
  /** The cursor width */
  cursorWidth: 2,
  /** Render the waveform with bars like this: ▁ ▂ ▇ ▃ ▅ ▂ */
  barWidth: 2,
  /** Spacing between bars in pixels */
  barGap: 3,
  /** Rounded borders for bars */
  barRadius: 10,
  /** A vertical scaling factor for the waveform */
  barHeight: NaN,
  /** Vertical bar alignment **/
  barAlign: '',
  /** Minimum pixels per second of audio (i.e. zoom level) */
  minPxPerSec: 20,
  /** Stretch the waveform to fill the container, true by default */
  fillParent: true,
  /** Audio URL */
  /** Whether to show default audio element controls */
  mediaControls: false,
  /** Play the audio on load */
  autoplay: false,
  /** Pass false to disable clicks on the waveform */
  interact: true,
  /** Allow to drag the cursor to seek to a new position */
  dragToSeek: false,
  /** Hide the scrollbar */
  hideScrollbar: false,
  /** Audio rate */
  audioRate: 1,
  /** Automatically scroll the container to keep the current position in viewport */
  autoScroll: true,
  /** If autoScroll is enabled, keep the cursor in the center of the waveform during playback */
  autoCenter: true,
  /** Decoding sample rate. Doesn't affect the playback. Defaults to 8000 */
  sampleRate: 8000,
}
export default {
  props: {
    content: { type: String, required: true },
  },
  data() {
    return {
      wavesurfer: null,
      currentTime: '00:00',
      duration: '00:01',
      onVolume: true,
      volumeValue: 1,
      playing: true,
      loading: true
    };
  },
  watch: {
    content(val){
        if(val){
            this.initWaveSurfer()
        }else{
            this.loading = false
        }
    }
  },
  methods: {
    sliderChange(event){
        this.volumeValue = parseFloat(event.target.value);
        if(this.volumeValue){
            this.onVolume = true
        }
        this.waveSurfer.setVolume(this.volumeValue)
    },
    offVolume(){
        this.onVolume = !this.onVolume
        if(this.onVolume){
            this.volumeValue = 1
            this.waveSurfer.setVolume(1)
        }else{
            this.volumeValue = 0
            this.waveSurfer.setVolume(0)
        }
    },
    skipBack(){
        this.waveSurfer.skip(-2)
    },
    skipForward(){
        this.waveSurfer.skip(2)
    },
    playPause(){
        this.waveSurfer.playPause()
    },
    format_time(seconds){
        const hours = Math.floor(seconds / 3600);
        const minutes = Math.floor((seconds % 3600) / 60);
        const seconds_remainder = Math.round(seconds) % 60;
        const padded_minutes = `${minutes < 10 ? "0" : ""}${minutes}`;
        const padded_seconds = `${
            seconds_remainder < 10 ? "0" : ""
        }${seconds_remainder}`;

        if (hours > 0) {
            return `${hours}:${padded_minutes}:${padded_seconds}`;
        }
        return `${minutes}:${padded_seconds}`;
    },
    download(){
        if(this.content){
            const a = document.createElement('a')
            a.href = this.content
            a.style.display = 'none'
            a.download = 'audio.wav';
            document.body.appendChild(a)
            a.click()
            document.body.removeChild(a)
            window.URL.revokeObjectURL(this.content) 
        }
    },

    initWaveSurfer() {
        this.waveSurfer = WaveSurfer.create({
            container: this.$refs.waveform,
            ...options
        })
        
        this.waveSurfer.on('click', () => {
            this.waveSurfer.play()
        })
        /** When audio starts loading */
        this.waveSurfer.on('load', (url) => {
        })
        /** During audio loading */
        this.waveSurfer.on('loading', (percent) => {
        })
        this.waveSurfer.on('decode', (duration) => {
            this.duration = this.format_time(duration)
            this.loading = false
        })
        /** When the audio starts playing */
        this.waveSurfer.on('play', () => {
            this.playing = false
        })
        /** When the audio pauses */
        this.waveSurfer.on('pause', () => {
            this.playing = true
        })
        /** When the audio finishes playing */
        this.waveSurfer.on('finish', () => {
        })
        /** On audio position change, fires continuously during playback */
        this.waveSurfer.on('timeupdate', (currentTime) => {
            this.currentTime = this.format_time(currentTime)
        })
        /** When the user drags the cursor */
        this.waveSurfer.on('drag', (relativeX) => {
        })
        this.waveSurfer.load(this.content);
    }
  },
  mounted() {
  },
  beforeDestroy() {
    this.waveSurfer.destroy();
  }
};
</script>

<style scoped lang="less">
.tts-container{
    position: relative;
    border: 1px solid rgba(229,231,235,1);
    border-radius: 10px 10px 10px 0px;
    background: #fff;
    width: 70%;
    height: 192px;
    line-height: 1.4;
    margin-bottom: 1rem;
    .tts-label{
        border-radius: calc( 8px - 1px) 0 calc( 8px - 1px) 0;
        border: 1px solid rgba(229,231,235,1);
        border-top: none;
        border-left: none;
        padding: 4px 12px;
        display: inline-block;
        background: #ffff;
        height:30px
    }
    .component-wrapper{
        padding: 12px;
        height: 100px;
        .waveform-container{
            display: flex;
            align-items: center;
            justify-content: center;
            width: 100%;
            #waveform{
                width: 100%;
                height: 100%;
                position: relative;
            }
        }
        .timestamps{
            display: flex;
            justify-content: space-between;
            align-items: center;
            width: 100%;
            padding: 4px 0;
            color:#9ca3af;
            font-size: 12px
        }
        .waveform-controls{
            display: grid;
            grid-template-columns: 1fr 1fr 1fr;
            margin-top: 5px;
            align-items: center;
            position: relative;
        }
        .control-wrapper{
            display: flex;
            justify-self: self-start;
            align-items: center;
            justify-content: space-between;
            .volume-icon{
                position: relative;
                display: flex;
                justify-content: center;
                margin-right: 10px;
                margin-left: 6px;
                font-size: 16px;
                color:#888
            }
            .volume-slider{
                background: #0191ff;
                width:80px;
                height: 4px;
                border-radius: 15px;
                cursor: pointer;
                outline: none;
            }
        }
        .play-pause-icon{
            display: flex;
            justify-self: center;
            font-size: 20px;
            color:#888;
            div{
                cursor: pointer;
            }
            .skip-left-icon{
                margin: 0 10px;
            }

            .skip-right-icon{
                margin: 0 10px;
            }
        }
        .download-icon{
            width: 24px;
            height: 24px;
            border-radius: 5px;
            background-color: #ffffff;
            border: 1px solid #e1e3e6;
            display: flex;
            justify-content: center;
            align-items: center;
            margin-left: auto;
            cursor: pointer;
        }
    }
}
</style>