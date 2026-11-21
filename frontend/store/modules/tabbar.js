export default {
	state: {
		list: [{
				text: 'Art',
				icon: 'home',
				inactiveIcon: '/static/tab_icon/art.png',
				activeIcon: '/static/tab_icon/art_fill.png',
				url: '/pages/index/index',
			},
			{
				text: 'Market',
				icon: 'account',
				inactiveIcon: '/static/tab_icon/market.png',
				activeIcon: '/static/tab_icon/market_fill.png',
				url: '/pages/market/market',
			},
			{
				text: 'Mine',
				icon: 'star',
				inactiveIcon: '/static/tab_icon/mine.png',
				activeIcon: '/static/tab_icon/mine_fill.png',
				url: '/pages/mine/mine',
			}
		],
		active: uni.getStorageSync('tabbarActive') || 0
	},
	mutations: {
		changeList(state, list) {
			state.list = list;
		},
		changeActiveIndex(state, index) {
			uni.setStorageSync('tabbarActive', index);
			state.active = index;
		}
	},
};
