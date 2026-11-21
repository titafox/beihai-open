/**
 * 设置语言类型： 默认为中文
 */
var i18nLanguage = "en";
// console.log(i18nLanguage.substr(0, 2))
/*
设置一下网站支持的语言种类
 */
var webLanguage = ['cn', 'en', 'vi'];

/*页面执行加载执行*/
$(function() {
	/*执行I18n翻译*/
	execI18n();

	/*将语言选择默认选中缓存中的值*/
	// $("#language option[value=" + i18nLanguage + "]").attr("selected", true);

	/* 选择语言 */
	// $("#language").on('change', function() {
	// 	var language = $(this).children('option:selected').val()
	// 	console.log(language);
	// 	getCookie("userLanguage", language, {
	// 		expires: 30,
	// 		path: '/'
	// 	});
	// 	location.reload();
	// });
});

/**
 * 执行页面i18n方法
 * @return
 */

var execI18n = function() {
	/*
	首先获取用户浏览器设备之前选择过的语言类型
	 */
	if (getCookie("userLanguage") &&getCookie("userLanguage") != 'undefined') {
		i18nLanguage = getCookie("userLanguage");
	} else {
		// 获取浏览器语言
		var navLanguage = getNavLanguage();
		if (navLanguage) {
			// 判断是否在网站支持语言数组里
			var charSize = $.inArray(navLanguage, webLanguage);
			if (charSize > -1) {
				console.log(navLanguage)
				i18nLanguage = navLanguage;
				// 存到缓存中
				getCookie("userLanguage", navLanguage);
			};
		} else {
			console.log("not navigator");
			return false;
		}
	}

	/* 需要引入 i18n 文件*/
	if ($.i18n == undefined) {
		console.log("请引入i18n js 文件")
		return false;
	};

	/*
	这里需要进行i18n的翻译
	 */
	jQuery.i18n.properties({
		name: 'strings', //资源文件名称
		path: '/static/web/locales/', //资源文件路径
		mode: 'map', //用Map的方式使用资源文件中的值
		language: i18nLanguage.substr(0, 2),
		callback: function() { //加载成功后设置显示内容 赋值
		
			/*以下为赋值方法*/
			// 普通元素赋值
			let insertEle = $('.i18n');
			insertEle.each(function() {
				$(this).html($.i18n.prop($(this).attr('name')));
				
			});
			
			//搜索占位文字赋值
			// console.log('input i18n')
			let insertInputEle = $(".i18n-input");
			insertInputEle.each(function(index, obj) {
				let selectAttr = $(this).attr('select-attr');
				if (!selectAttr) {
					selectAttr = "value";
				};
				$(this).attr(selectAttr, $.i18n.prop($(this).attr('select-name')));
			});
		}
	});
}

/**
 * cookie操作
 */
var getCookie = function(name, value, options) {
	
	if(!options){
		options = {}
		options.expires = 30;
		options.path =  '/'
	}
	
	
	
	if (typeof value != 'undefined') { // name and value given, set cookie typeof value != 'undefined'
		options = options || {};
		if (value === null) {
			value = '';
			options.expires = -1;
		}
		var expires = '';
		if (options.expires && (typeof options.expires == 'number' || options.expires.toUTCString)) {
			var date;
			if (typeof options.expires == 'number') {
				date = new Date();
				date.setTime(date.getTime() + (options.expires * 24 * 60 * 60 * 1000));
				// console.log(date)
			} else {
				date = options.expires;
			}
			expires = '; expires=' + date.toUTCString(); // use expires attribute, max-age is not supported by IE
		}
		var path = options.path ? '; path=' + options.path : '';
		var domain = options.domain ? '; domain=' + options.domain : '';
		var s = [cookie, expires, path, domain, secure].join('');
		var secure = options.secure ? '; secure' : '';
		var c = [name, '=', encodeURIComponent(value)].join('');
		var cookie = [c, expires, path, domain, secure].join('')
		// console.log(expires)
		document.cookie = cookie;
	} else { // only name given, get cookie
		var cookieValue = null;
		if (document.cookie && document.cookie != '') {
			var cookies = document.cookie.split(';');
			for (var i = 0; i < cookies.length; i++) {
				var cookie = jQuery.trim(cookies[i]);
				// Does this cookie string begin with the name we want?
				if (cookie.substring(0, name.length + 1) == (name + '=')) {
					cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
					break;
				}
			}
		}
		return cookieValue;
	}
};

/**
 * 获取浏览器语言类型
 * @return {string} 浏览器国家语言
 */

var getNavLanguage = function() {
	if (navigator.appName == "Netscape") {
		var navLanguage = navigator.language;
		return navLanguage.substr(0, 5).toLowerCase();
	}
	return false;
}
