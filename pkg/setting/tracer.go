package setting

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName, // 对其发起请求的的调用链，叫什么服务
		Sampler: &config.SamplerConfig{
			/*"const"	0或1	采样器始终对所有 tracer 做出相同的决定；要么全部采样，要么全部不采样
			"probabilistic"	0.0~1.0	采样器做出随机采样决策，Param 为采样概率
			"ratelimiting"	N	采样器一定的恒定速率对tracer进行采样，Param=2.0，则限制每秒采集2条
			"remote"	无	采样器请咨询Jaeger代理以获取在当前服务中使用的适当采样策略。*/
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			/*QUEUESIZE，设置队列大小，存储采样的 span 信息，队列满了后一次性发送到 jaeger 后端；defaultQueueSize 默认为 100；
			BufferFlushInterval 强制清空、推送队列时间，对于流量不高的程序，队列可能长时间不能满，那么设置这个时间，超时可以自动推送一次。对于高并发的情况，一般队列很快就会满的，满了后也会自动推送。默认为1秒。
			LogSpans 是否把 Log 也推送，span 中可以携带一些日志信息。
			LocalAgentHostPort 要推送到的 Jaeger agent，默认端口 6831，是 Jaeger 接收压缩格式的 thrift 协议的数据端口。
			CollectorEndpoint 要推送到的 Jaeger Collector，用 Collector 就不用 agent 了。*/
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer() // 根据配置初始化 tracer对象
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
