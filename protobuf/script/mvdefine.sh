if [ $1 == 'define' ]
then 
    mv grpc/define/github.com/lzl-here/mini_video_proto/sdk/protocol/grpc/define/*.pb.go $2
    rm -rf grpc/define/github.com/lzl-here/mini_video_proto
else
    mv $2/$1/*.pb.gw.go $2
	rm -rf $2/$1
fi

