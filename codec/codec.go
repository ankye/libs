package codec

//Codec 编解码接口定义
type Codec interface {
	Name() string
	Encode(obj interface{}) ([]byte, error)
	Decode(data []byte, obj interface{}) error
}
