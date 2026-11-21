<?php

namespace app\admin\model\tbl;

use think\Model;


class Handsel extends Model
{

    

    

    // 表名
    protected $table = 'tbl_handsel';
    
    // 自动写入时间戳字段
    protected $autoWriteTimestamp = false;

    // 定义时间戳字段名
    protected $createTime = false;
    protected $updateTime = false;
    protected $deleteTime = false;

    // 追加属性
    protected $append = [
        'typ_text',
        'crete_time_text'
    ];
    

    
    public function getTypList()
    {
        return ['0' => __('Typ 0'), '1' => __('Typ 1')];
    }


    public function getTypTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['typ']) ? $data['typ'] : '');
        $list = $this->getTypList();
        return isset($list[$value]) ? $list[$value] : '';
    }


    public function getCreteTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['crete_time']) ? $data['crete_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }

    protected function setCreteTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }


    public function works()
    {
        return $this->belongsTo('Works', 'w_id', 'id', [], 'LEFT')->setEagerlyType(0);
    }


    public function account()
    {
        return $this->belongsTo('Account', 'account_id', 'id', [], 'LEFT')->setEagerlyType(0);
    }
}
