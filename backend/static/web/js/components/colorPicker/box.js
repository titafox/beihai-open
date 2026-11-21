Vue.component('box', {
	template: ` 
	<div class="color-type">
        <span class="name">
            {{ name }}
        </span>
        <input
            v-model="modelColor"
            class="value"
        >
    </div>`,
	props: {
		name: {
			type: String,
			default: ''
		},
		color: {
			type: String,
			default: ''
		}
	},
	computed: {
		modelColor: {
			get() {
				return this.color
			},
			set(val) {
				this.$emit('inputColor', val)
			}
		}
	}
})
