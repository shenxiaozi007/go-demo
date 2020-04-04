<?php
class test{
    public function getMoneyFunc() {
        $rmb = 1;
        $func = function () use (&$rmb) {
            echo "闭包里面---" . $rmb . "\n";
            $rmb++;
        };
        $func2 = function () use (&$rmb) {
            echo "闭包里面===" . $rmb . "\n";
            $rmb++;
        };
        $func();
        $func2();
        return $func;
    }
}

$test = new test();
$func = $test->getMoneyFunc();
$func();