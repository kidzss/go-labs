<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: tabletmanagerdata.proto

namespace Vitess\Proto\Tabletmanagerdata {

  class PreflightSchemaRequest extends \DrSlump\Protobuf\Message {

    /**  @var string[]  */
    public $changes = array();
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'tabletmanagerdata.PreflightSchemaRequest');

      // REPEATED STRING changes = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "changes";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <changes> has a value
     *
     * @return boolean
     */
    public function hasChanges(){
      return $this->_has(1);
    }
    
    /**
     * Clear <changes> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\PreflightSchemaRequest
     */
    public function clearChanges(){
      return $this->_clear(1);
    }
    
    /**
     * Get <changes> value
     *
     * @param int $idx
     * @return string
     */
    public function getChanges($idx = NULL){
      return $this->_get(1, $idx);
    }
    
    /**
     * Set <changes> value
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\PreflightSchemaRequest
     */
    public function setChanges( $value, $idx = NULL){
      return $this->_set(1, $value, $idx);
    }
    
    /**
     * Get all elements of <changes>
     *
     * @return string[]
     */
    public function getChangesList(){
     return $this->_get(1);
    }
    
    /**
     * Add a new element to <changes>
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\PreflightSchemaRequest
     */
    public function addChanges( $value){
     return $this->_add(1, $value);
    }
  }
}

