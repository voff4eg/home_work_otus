package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := doInOut(in, done)
	for _, s := range stages {
		if s == nil {
			panic("stage should be not nil")
		}
		out = s(doInOut(out, done))
	}

	return out
}

func doInOut(in In, do Out) Out {
	out := make(Bi)
	go func() {
		defer close(out)

		for {
			select {
			case <-do:
				return
			case k, v := <-in:
				if !v {
					return
				}
				out <- k
			}
		}
	}()

	return out
}
