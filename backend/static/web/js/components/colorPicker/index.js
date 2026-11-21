Vue.component('color-picker',{
	template:`
		<div class="color-picker" :class="{ light: isLightTheme }" :style="{ width: totalWidth + 'px' }">
			<div class="color-set">
		      <saturation ref="saturation" :color="rgbString" :hsv="hsv" :size="hueHeight"
		        @selectSaturation="selectSaturation" />
		      <hue ref="hue" :hsv="hsv" :width="hueWidth" :height="hueHeight" @selectHue="selectHue" />
		      <alpha ref="alpha" :color="rgbString" :rgba="rgba" :width="hueWidth" :height="hueHeight"
		        @selectAlpha="selectAlpha" />
		    </div>
		    <div :style="{ height: previewHeight + 'px' }" class="color-show">
		      <preview :color="rgbaString" :width="previewWidth" :height="previewHeight" />
		      <sucker v-if="!suckerHide" :sucker-canvas="suckerCanvas" :sucker-area="suckerArea" @openSucker="openSucker"
		        @selectSucker="selectSucker" />
		    </div>
		    <box name=" HEX" :color="modelHex" @inputColor="inputHex" />
		    <box name="RGBA" :color="modelRgba" @inputColor="inputRgba" />
		    <colors :color="rgbaString" :colors-default="colorsDefault" :colors-history-key="colorsHistoryKey"
		      @selectColor="selectColor" @newColor="newColor" :color-count="colorCount" />
		  </div>
	`,
	props: {
	      color: {
	        type: String,
	        default: '#000000'
	      },
	      theme: {
	        type: String,
	        default: 'dark'
	      },
	      suckerHide: {
	        type: Boolean,
	        default: true
	      },
	      suckerCanvas: {
	        type: null, // HTMLCanvasElement
	        default: null
	      },
	      suckerArea: {
	        type: Array,
	        default: () => []
	      },
	      colorsDefault: {
	        type: Array,
	        default: () => [
	          '#000000', '#FFFFFF', '#FF1900', '#F47365', '#FFB243', '#FFE623', '#6EFF2A', '#1BC7B1',
	          '#00BEFF', '#2E81FF', '#5D61FF', '#FF89CF', '#FC3CAD', '#BF3DCE', '#8E00A7', 'rgba(0,0,0,0)'
	        ]
	      },
	      colorsHistoryKey: {
	        type: String,
	        default: 'vue-colorpicker-history'
	      }
	    },
	    data() {
	      return {
	        hueWidth: 15,
	        hueHeight: 152,
	        previewHeight: 30,
	        modelRgba: '',
	        modelHex: '',
	        r: 0,
	        g: 0,
	        b: 0,
	        a: 1,
	        h: 0,
	        s: 0,
	        v: 0,
	        colorCount: 0
	      }
	    },
	    computed: {
	      isLightTheme() {
	        return this.theme === 'light'
	      },
	      totalWidth() {
	        return this.hueHeight + (this.hueWidth + 8) * 2
	      },
	      previewWidth() {
	        return this.totalWidth - (this.suckerHide ? 0 : this.previewHeight)
	      },
	      rgba() {
	        return {
	          r: this.r,
	          g: this.g,
	          b: this.b,
	          a: this.a
	        }
	      },
	      hsv() {
			
	        return {
	          h: this.h,
	          s: this.s,
	          v: this.v
	        }
	      },
	      rgbString() {
	        return `rgb(${this.r}, ${this.g}, ${this.b})`
	      },
	      rgbaStringShort() {
	        return `${this.r}, ${this.g}, ${this.b}, ${this.a}`
	      },
	      rgbaString() {
	        return `rgba(${this.rgbaStringShort})`
	      },
	      hexString() {
	        return rgb2hex(this.rgba, true)
	      }
	    },
	    created() {
			
			Object.assign(this, setColorValue(this.color))
			this.setText()
	
	      // 避免初始化时，也会触发changeColor事件
	      this.$watch('rgba', () => {
	        this.$emit('changecolor', {
	          rgba: this.rgba,
	          hsv: this.hsv,
	          hex: this.modelHex
	        })
	      })
	    },
	    methods: {
	      selectSaturation(color) {
			
	        const { r, g, b, h, s, v, a } = setColorValue(color)
			
			
			console.log(a)
			
	        Object.assign(this, { r, g, b, h, s, v, a })
			
	        this.setText()
	      },
	      selectHue(color) {
	        const { r, g, b, h, s, v } = setColorValue(color)
	        Object.assign(this, { r, g, b, h, s, v })
	        this.setText()
	        this.$nextTick(() => {
	          this.$refs.saturation.renderColor()
	          this.$refs.saturation.renderSlide()
	        })
	      },
	      selectAlpha(a) {
	        this.a = a
	        this.setText()
	      },
	      inputHex(color) {
	        const { r, g, b, a, h, s, v } = setColorValue(color)
	        Object.assign(this, { r, g, b, a, h, s, v })
	        this.modelHex = color
	        this.modelRgba = this.rgbaStringShort
	        this.$nextTick(() => {
	          this.$refs.saturation.renderColor()
	          this.$refs.saturation.renderSlide()
	          this.$refs.hue.renderSlide()
	        })
	      },
	      inputRgba(color) {
	        const { r, g, b, a, h, s, v } = setColorValue(color)
	        Object.assign(this, { r, g, b, a, h, s, v })
	        this.modelHex = this.hexString
	        this.modelRgba = color
	        this.$nextTick(() => {
	          this.$refs.saturation.renderColor()
	          this.$refs.saturation.renderSlide()
	          this.$refs.hue.renderSlide()
	        })
	      },
	      setText() {
			
	        this.modelHex = this.hexString
	        this.modelRgba = this.rgbaStringShort
			
	      },
	      openSucker(isOpen) {
				console.log('openSucker index')
				this.$emit('opensucker', isOpen)
	      },
	      selectSucker(color) {
	        const { r, g, b, a, h, s, v } = setColorValue(color)
	        this.r = r;
	        this.g = g;
	        this.b = b;
	        this.a = a;
	        this.h = h;
	        this.s = s;
	        this.v = v;
	        // Object.assign(this, { r, g, b, a, h, s, v })
	
	        this.setText()
	        this.$nextTick(() => {
	          this.$refs.saturation.renderColor()
	          this.$refs.saturation.renderSlide()
	          this.$refs.hue.renderSlide()
	        })
	      },
	      selectColor(color) {
	        const { r, g, b, a, h, s, v } = setColorValue(color)
	        Object.assign(this, { r, g, b, a, h, s, v })
	
	        this.setText()
	        this.$nextTick(() => {
	          this.$refs.saturation.renderColor()
	          this.$refs.saturation.renderSlide()
	          this.$refs.hue.renderSlide()
	        })
	      },
	      newColor() {
	        let colorsHistoryKey = JSON.parse(localStorage.getItem(this.colorsHistoryKey));
	        const color = `rgba(${this.rgbaStringShort})`;
			if(!colorsHistoryKey) colorsHistoryKey = []
	        if (colorsHistoryKey.indexOf(color) == -1) {
	          colorsHistoryKey.unshift(color);
	          localStorage.setItem(this.colorsHistoryKey, JSON.stringify(colorsHistoryKey))
	          this.colorCount++;
	        }
	
	      }
	    }
})