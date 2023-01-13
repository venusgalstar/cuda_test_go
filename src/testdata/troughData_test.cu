extern "C" 
__global__ void throughData(double* A, double* B, int* C, int size) {
    int block = blockIdx.x + blockIdx.y * gridDim.x + gridDim.x * gridDim.y * blockIdx.z;
    int index = block * (blockDim.x * blockDim.y * blockDim.z) + (threadIdx.z * (blockDim.x * blockDim.y)) + (threadIdx.y * blockDim.x) + threadIdx.x;
    if(index >= size) return;

    if( A[index] >= B[index] )
        C[index] = 0;
    else C[index] = 1;
}