<?php

function unique_rand($min, $max, $num) {
  $count = 0;
  $return = array();
  while ($count < $num) {
    $return[] = mt_rand($min, $max);
    $return = array_flip(array_flip($return));
    $count = count($return);
  }
  //打乱数组，重新赋予数组新的下标
  shuffle($return);
  return $return;
}

//生成10个1到100范围内的不重复随机数
$s = time();
$arr = unique_rand(1, 10, 10);
print_r($arr);

echo time()-$s;

