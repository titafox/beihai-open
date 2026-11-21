export default {
	namespaced: true,
	state: {
		b_count:0,
		m_count:0
	},
	mutations: {
		buyStateHandler(state,count){
			state.b_count = state.b_count+count;
		},
		mintStateHandler(state){
			state.m_count = state.m_count+count;
		}
	},
};
