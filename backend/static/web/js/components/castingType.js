//铸件类型
Vue.component('casting-type', {
	template: `
		<div class="casting-type">
			<div class="tab-content">
				<div class="tab-pane fade show active" role="tabpanel">
					<div class="row row--grid">
						<div class="col-12 col-sm-6 col-lg-4 casting-type-container" v-for="(item, i) in typelist" :key="i" @click="$emit('viewclick', (i+1))">
							<div class="card">
								<a href="javascript:void(0)" class="card__cover">
									<img :src="item.logo" alt="">
								</a>
								<div class="card__author">
									<a>{{item.num}}</a>
								</div>
								<h3 class="card__title">
									<a href="javascript:void(0)">{{item.num}}</a>
								</h3>
							</div>
							<h2 class="name i18n" :name="item.name"></h2>
						</div>
					</div>
				</div>
			</div>
		</div>
	`,
	data: function() {
		return {
			typelist:[
				{
					name: 'create-singleNFT',
					logo: 'img/components/single.png',
					num: '1 / 1'
				},
				{
					name: 'create-seriesNFTs',
					logo: 'img/components/series.png',
					num: '1 / 1000'
				},
				{
					name: 'create-importNFTs',
					logo: 'img/components/import.png',
					num: '1 / 1000'
				}
			]
		}
	},
	methods: {
		
	}
})


