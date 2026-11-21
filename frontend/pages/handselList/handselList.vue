<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="我的寄售" :placeholder="true" class="nav_far"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-gap height="44"></u-gap>

		<u-empty v-if="dblist.length === 0" mode="favor" text="啥都还没买过"></u-empty>
		<ShopCard :list="dblist"></ShopCard>


	</view>
</template>

<script>
	import {
		nftsells
	} from '@/api/index.js';
	import ShopCard from '../../components/ShopCard2/ShopCard.vue'


	export default {
		components: {

			ShopCard
		},

		data() {
			return {

				dblist: [],

			};
		},
		onShow() {
			this.getTotalList();
		},
		methods: {
			getTotalList() {

				nftsells({

				}).then(res => {
					console.log(res);
					this.dblist = res;
				}).catch(() => {

				});
			},


		}
	};
</script>
<style scoped>

</style>
