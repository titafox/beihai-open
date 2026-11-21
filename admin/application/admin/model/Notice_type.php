<?php

namespace app\admin\model;

use think\Model;


class Notice_type extends Model
{

    

    

    // 表名
    protected $table = 'tbl_notice_type';
    
    // 自动写入时间戳字段
    protected $autoWriteTimestamp = false;

    // 定义时间戳字段名
    protected $createTime = false;
    protected $updateTime = false;
    protected $deleteTime = false;

    // 追加属性
    protected $append = [
        'display_text'
    ];
    

    
    public function getDisplayList()
    {
        return ['1' => __('Display 1'), '0' => __('Display 0')];
    }


    public function getDisplayTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['display']) ? $data['display'] : '');
        $list = $this->getDisplayList();
        return isset($list[$value]) ? $list[$value] : '';
    }




}
