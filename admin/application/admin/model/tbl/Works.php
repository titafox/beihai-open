<?php

namespace app\admin\model\tbl;

use think\Model;


class Works extends Model
{

    

    

    // 表名
    protected $table = 'tbl_works';
    
    // 自动写入时间戳字段
    protected $autoWriteTimestamp = false;

    // 定义时间戳字段名
    protected $createTime = false;
    protected $updateTime = false;
    protected $deleteTime = false;

    // 追加属性
    protected $append = [
        'pub_time_text',
        'status_text',
        'create_time_text',
        'update_time_text',
        'sold_out_text',
        'sell_time_text',
        'typ_text',
        'is_synthetic_material',
    ];
    

    
    public function getStatusList()
    {
        return ['0' => __('Status 0'), '1' => __('Status 1')];
    }

    public function getSoldOutList()
    {
        return ['0' => __('Sold_out 0'), '1' => __('Sold_out 1')];
    }

    public function getTypList()
    {
      return ['1' => __('Typ 1'), '2' => __('Typ 2'),'3' => __('Typ 3')];
    }

  public function getIsSyntheticMaterialList()
  {
    return ['0' => __('Is_synthetic_material 0'), '1' => __('Is_synthetic_material 1')];
  }


    public function getPubTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['pub_time']) ? $data['pub_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }


    public function getStatusTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['status']) ? $data['status'] : '');
        $list = $this->getStatusList();
        return isset($list[$value]) ? $list[$value] : '';
    }


    public function getCreateTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['create_time']) ? $data['create_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }


    public function getUpdateTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['update_time']) ? $data['update_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }


    public function getSoldOutTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['sold_out']) ? $data['sold_out'] : '');
        $list = $this->getSoldOutList();
        return isset($list[$value]) ? $list[$value] : '';
    }

  public function getTypTextAttr($value, $data)
  {
    $value = $value ? $value : (isset($data['typ']) ? $data['typ'] : '');
    $list = $this->getTypList();
    return isset($list[$value]) ? $list[$value] : '';
  }

  public function getIsSyntheticMaterTextAttr($value, $data)
  {
    $value = $value ? $value : (isset($data['is_synthetic_material']) ? $data['is_synthetic_material'] : '');
    $list = $this->getIsSyntheticMaterialList();
    return isset($list[$value]) ? $list[$value] : '';
  }


    public function getSellTimeTextAttr($value, $data)
    {
        $value = $value ? $value : (isset($data['sell_time']) ? $data['sell_time'] : '');
        return is_numeric($value) ? date("Y-m-d H:i:s", $value) : $value;
    }

    protected function setPubTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }

    protected function setCreateTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }

    protected function setUpdateTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }

    protected function setSellTimeAttr($value)
    {
        return $value === '' ? null : ($value && !is_numeric($value) ? strtotime($value) : $value);
    }


    public function ips()
    {
        return $this->belongsTo('Ips', 'author_id', 'id', [], 'LEFT')->setEagerlyType(0);
    }
}
