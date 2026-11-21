export function stop(){
  var box=function(e){passive: false ;};
  document.body.style.overflow='hidden';
  document.addEventListener("touchmove",box,false);
}
// 允许页面滚动
export function move(){
  var box=function(e){passive: false };
  document.body.style.overflow=''; // 出现滚动条
  document.removeEventListener("touchmove",box,false);
}
