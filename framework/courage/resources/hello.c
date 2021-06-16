import "error.IOError"
import "io"

int main()
{
    println("hello world!!!");

    int[] arr={1,2,3,4,5};

    for(int a:arr)
    {
        println(a);
    }

    int a=1*2+(3+6)/2;

    try
    {
       String text,long len=readFile("C:\\Users\\Admin\\Desktop\\hello.txt");
       println("文件内容="+text);
       println("文件长度="+len);
    }catch(IOError e)
    {
       println("读取文件发送异常="+e.getMessage());
    }
    return 0;
}

String,long readFile(String path) throws IOError
{
    File file=io.Open(path)
    return io.ReadAll(file)
}
