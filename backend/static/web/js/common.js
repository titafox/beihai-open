function showModal(options){
	
	$('.error-modal').find('.modal-title').html('');
	$('.error-modal').find('p').siblings('img').remove();
	console.log(options.type)
	
	if(options.type == 'success'){
		$('.error-modal').find('p').before(
			"<img src='/static/web/img/common/success.svg' />"
		);
	} else {
		$('.error-modal').find('p').before(
			"<img src='/static/web/img/common/fail.svg' />"
		);
	}
	$('.error-modal').find('p').html(options.message);
	$('.error-modal').removeClass('fadeOut').addClass('fadeIn');
	
} 

function hideModal(){
	$('.modal').removeClass('fadeIn').addClass('fadeOut');
} 
function showToast(options){
	hideToast()
	$('.toast').find('.toast-header').html('');
	$('.toast').find('.toast-body').html('');
	options = Object.assign({
	    type: 'success',// success error loading info
	    message: '',
	    duration: 0
	}, options);
	console.log(options)
	
	
	if(options.type == 'success'){
		$('.toast').find('.toast-header').append(
			"<img src='/static/web/img/common/success.svg' />"
		);
	} else if(options.type == 'loading') {
		$('.toast').find('.toast-header').append(
			"<img src='/static/web/img/common/loading.svg' class='loading'/>"
		);
	} else if(options.type == 'info'){
		$('.toast').find('.toast-header').append(
			"<img src='/static/web/img/common/infomation.svg' />"
		);
	} else {
		$('.toast').find('.toast-header').append(
			"<img src='/static/web/img/common/fail.svg' />"
		);
	}
	
	if(options.title){
		$('.toast').find('.toast-header').find('img').after('<p>'+options.title+'</p>');
	}
	
	
	
	$('.toast').find('.toast-body').html(options.message);
	$('.toast').removeClass('fadeOut').addClass('fadeIn');
	
	if(options.duration){
		setTimeout(()=>{
			if(options.callback){
				options.callback()
			}
			hideToast()
		}, options.duration)//
	}
	
	
} 

function hideToast(){
	$('.toast').removeClass('fadeIn').addClass('fadeOut');
}

function getQueryString(name){
	var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
	var r = location.search.substr(1).match(reg);
	if (r != null) return unescape(decodeURI(r[2])); return null;
}

//时间戳转时间
function time(timer){
  var arrTimestamp = (timer + '').split('');
	for (var start = 0; start < 13; start++) {
		if (!arrTimestamp[start]) {
			arrTimestamp[start] = '0';
		}
	}
	var Etime = arrTimestamp.join('') * 1; //参数时间
	var Etimer = new Date(Etime)

	var Eyear = Etimer.getFullYear(); //取得4位数的年份
	var Emonth = Etimer.getMonth() + 1 < 10 ? '0' + (Etimer.getMonth() + 1) : Etimer.getMonth() +
		1; //取得日期中的月份，其中0表示1月，11表示12月
	var Edate = Etimer.getDate() < 10 ? '0' + Etimer.getDate() : Etimer.getDate(); //返回日期月份中的天数（1到31）
	var Ehour = Etimer.getHours() < 10 ? '0' + Etimer.getHours() : Etimer.getHours(); //返回日期中的小时数（0到23）
	var Eminute = Etimer.getMinutes() < 10 ? '0' + Etimer.getMinutes() : Etimer.getMinutes(); //返回日期中的分钟数（0到59）
	var Esecond = Etimer.getSeconds() < 10 ? '0' + Etimer.getSeconds() : Etimer.getSeconds(); //返回日期中的分钟数（0到59）
	return Eyear + "-" + Emonth + "-" + Edate + " " + Ehour + ":" + Eminute + ':' + Esecond;
}

// 数字三位用逗号隔开
function readNumber(number){
  if (!number) return '0'
  var intPart = Number(number).toFixed(0) // 获取整数部分
  var intPartFormat = intPart.toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,') // 将整数部分逢三一断
  var floatPart = '.00' // 预定义小数部分
  var value2Array = number.toString().split('.')
  // =2表示数据有小数位 
  if (value2Array.length === 2) {
    floatPart = value2Array[1].toString() // 拿到小数部分
    if (floatPart.length === 1) {
    // 补0
      return intPartFormat + '.' + floatPart + '0'
    } else {
      return intPartFormat + '.' + floatPart
    }
  } else {
    return intPartFormat
  }
}

function formatWallet(wallet){
	if(wallet){
		return wallet.substring(0,10) + '...' + wallet.substring(wallet.length - 6, wallet.length)
	} else {
		return ''
	}
}


//TODO components 里面的公用方法

