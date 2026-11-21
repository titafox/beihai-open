<?php

namespace app\admin\model\tbl;

use think\Model;


class Snapshot extends Model
{

    

    

    // 表名
    protected $table = 'tbl_snapshot';
    
    // 自动写入时间戳字段
    protected $autoWriteTimestamp = false;

    // 定义时间戳字段名
    protected $createTime = false;
    protected $updateTime = false;
    protected $deleteTime = false;

    // 追加属性
    protected $append = [
        'start_time_text',
        'snapshot_mod_text',
        'update_time_text'
    ];

  public function getStatusList()
  {
    return ['0' => __('Status 0'), '1' => __('Status 1'),'2' => __('Status 2')];
  }


  public function getStatusTextAttr($value, $data)
  {
    $value = $value ? $value : (isset($data['status']) ? $data['status'] : '');
    $list = $this->getStatusList();
    return isset($list[$value]) ? $list[$value] : '';
  }
    

    
    public function getSnapshotModList()
    {
        return ['0' => __('Snapshot_mod 0'), '1' => __('Snapshot_mod 1')];
    }


    public function getStartTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['start_time']) ? $data['start_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }


    public function getSnapshotModTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['snapshot_mod']) ? $data['snapshot_mod'] : '');
        $list = $this->getSnapshotModList();
        return isset($list[$value]) ? $list[$value] : '';
    }


    public function getUpdateTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['update_time']) ? $data['update_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }

    protected function setStartTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }

    protected function setUpdateTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }


    public function works()
    {
        return $this->belongsTo('Works', 'a_id', 'id', [], 'LEFT')->setEagerlyType(0);
    }

  public function bworks()
  {
    return $this->belongsTo('Works', 'b_id', 'id', [], 'LEFT')->setEagerlyType(0);
  }
}
