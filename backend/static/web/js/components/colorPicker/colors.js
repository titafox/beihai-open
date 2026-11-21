Vue.component('colors', {
	template: ` 
		<div>
			<ul class="colors">
			  <li class="item" @click="newColor">
				<div class="alpha" style="display: flex;
				align-items: center;
				justify-content: center;
				border: 1px solid #cfcfcf;color: #cfcfcf;">
				  <span>+</span>
				</div>
			  </li>
			  <li v-for="item in colorsDefault" :key="item" class="item" @click="selectColor(item)">
				<div :style="{background:'url('+imgAlphaBase64+')'}" class="alpha" />
				<div :style="{ background: item }" class="color" />
			  </li>
			</ul>
			<ul v-if="colorsHistory.length" class="colors history">
			  <li v-for="item in colorsHistory" :key="item" class="item" @click="selectColor(item)">
				<div :style="{background:'url('+imgAlphaBase64+')'}" class="alpha" />
				<div :style="{ background: item }" class="color" />
			  </li>
			</ul>
		</div>`,
	 props: {
	      color: {
	        type: String,
	        default: '#000000'
	      },
	      colorsDefault: {
	        type: Array,
	        default: () => []
	      },
	      colorsHistoryKey: {
	        type: String,
	        default: ''
	      },
	      colorCount: {
	        type: Number,
	        default: 0
	      }
	    },
	    data() {
	      return {
	        imgAlphaBase64: '',
	        colorsHistory: JSON.parse(localStorage.getItem(this.colorsHistoryKey)) || []
	      }
	    },
	    watch: {
	      colorCount(newval) {
	        this.colorsHistory = JSON.parse(localStorage.getItem(this.colorsHistoryKey))
	      }
	    },
	    created() {
		
	      this.imgAlphaBase64 = createAlphaSquare(4).toDataURL()
	    },
	    destroyed() {
	      this.setColorsHistory(this.color)
	    },
	    methods: {
	      selectColor(color) {
	        this.$emit('selectColor', color)
	      },
	      newColor() {
	        this.$emit('newColor');
	
	      },
	      setColorsHistory(color) {
	        if (!color) {
	          return
	        }
	        const colors = this.colorsHistory
	        const index = colors.indexOf(color)
	        if (index >= 0) {
	          colors.splice(index, 1)
	        }
	        if (colors.length >= 8) {
	          colors.length = 7
	        }
	        colors.unshift(color)
	        this.colorsHistory = colors
	        localStorage.setItem(this.colorsHistoryKey, JSON.stringify(colors))
	      }
	}
	
})
