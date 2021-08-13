package main


// OKな例
type Stringer interface {
    // そもそもunionsがないので問題なし
    String() string 
}

type Number interface {
    // unionsがあるが、termであるintとfloatはいずれもnon-interface型なので問題なし
    ~int | ~float64 
}

type C1 interface {
    // NumberはMethodを持たないInterfaceなので、unionsの項(term)になることができる。	
    Number | ~string
    // StringerはMethodを含むInterfaceだが、"stand-alone"で埋め込まれているのでOK
    Stringer
    m()
}

type C2 interface {
    // C1 はMethodをもつInterfaceだが、"stand-alone"で埋め込まれているのでOK
    C1
}

type NotAllowed interface {
	~int | Stringer
}

func main() {
}
