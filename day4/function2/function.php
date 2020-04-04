<?php
class test{
    public function getMoneyFunc() {
        $rmb = 1;
        $func = function () use (&$rmb) {
            echo "闭包里面" . $rmb . "\n";
            $rmb++;
        };
        echo "闭包外面" . $rmb . "\n";
        return $func;
    }
}

$test = new test();
$func = $test->getMoneyFunc();
$func();
$func();
$func2 = $test->getMoneyFunc();