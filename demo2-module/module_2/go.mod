module github.com/qiudeng7/golang-lab.git/demo2-module/module_2

go 1.25.3

replace github.com/qiudeng7/golang-lab.git/demo2-module/module_1 => ../module_1

require github.com/qiudeng7/golang-lab.git/demo2-module/module_1 v0.0.0-00010101000000-000000000000
