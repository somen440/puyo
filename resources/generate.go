package resources

//go:generate file2byteslice -package=images -input=./images/puyo.png -output=./images/puyo.go -var=Puyo_png
//go:generate file2byteslice -package=images -input=./images/stage.png -output=./images/stage.go -var=Stage_png
//go:generate gofmt -s -w .
