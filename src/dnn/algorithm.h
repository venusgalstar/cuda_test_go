extern cudnnAlgorithm_t makeConvFwdAlgo(cudnnConvolutionFwdAlgo_t algo);
extern cudnnAlgorithm_t makeConvBwdFilterAlgo(cudnnConvolutionBwdFilterAlgo_t algo);
extern cudnnAlgorithm_t makeConvBwdDataAlgo(cudnnConvolutionBwdDataAlgo_t algo);
extern cudnnAlgorithm_t makeRNNAlgo(cudnnRNNAlgo_t algo);
extern cudnnAlgorithm_t makeCTCLossAlgo(cudnnCTCLossAlgo_t algo);
