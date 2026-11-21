<?php

namespace app\admin\model\tbl;

use think\Model;


class Nfts extends Model
{

    

    

    // 表名
    protected $table = 'tbl_nfts';
    
    // 自动写入时间戳字段
    protected $autoWriteTimestamp = false;

    // 定义时间戳字段名
    protected $createTime = false;
    protected $updateTime = false;
    protected $deleteTime = false;

    // 追加属性
    protected $append = [
        'chain_status_text',
        'status_text',
        'update_time_text'
    ];
    

    
    public function getChainStatusList()
    {
        return ['0' => __('Chain_status 0'), '1' => __('Chain_status 1'), '2' => __('Chain_status 2')];
    }

    public function getStatusList()
    {
        return ['0' => __('Status 0'), '1' => __('Status 1'), '2' => __('Status 2'), '3' => __('Status 3')];
    }


    public function getChainStatusTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['chain_status']) ? $data['chain_status'] : '');
        $list = $this->getChainStatusList();
        return isset($list[$value]) ? $list[$value] : '';
    }


    public function getStatusTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['status']) ? $data['status'] : '');
        $list = $this->getStatusList();
        return isset($list[$value]) ? $list[$value] : '';
    }


    public function getUpdateTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['update_time']) ? $data['update_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }

    protected function setUpdateTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }


    public function account()
    {
        return $this->belongsTo('Account', 'owner_id', 'id', [], 'LEFT')->setEagerlyType(0);
    }


    public function works()
    {
        return $this->belongsTo('Works', 'wid', 'id', [], 'LEFT')->setEagerlyType(0);
    }
}
