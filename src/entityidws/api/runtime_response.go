package api

type RuntimeResponse struct {
   CpuCount         int    `json:"cpus"`
   GoRoutineCount   int    `json:"go_routines"`
   ObjectCount      uint64 `json:"heap_objects"`
   AllocatedMemory  uint64 `json:"allocated_mem"`
}

