// file RBFNet.go

package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
)


type RadialBasis struct { // 2 layer neural network
    learningRate float64
    k int
    centers []float64
    standardDevs []float64
    epochs int
    weights []float64
    bias float64
}
// initialize the model to start at a random hypothesis, then following some
// learning rate until reach a presumably ideal state (local optima)
func abs(n float64) float64 {
    if (n < 0) {
        n *= float64(-1);
    }
    return n;
}

func Constructor(learningRate float64, k int, epochs int) RadialBasis {
    newRBF := new(RadialBasis);
    newRBF.learningRate = learningRate;
    newRBF.k = k;
    newRBF.epochs = epochs // unit of passing data forward & backwards NN once
    newRBF.weights = []float64{};
    // idea is to update weights of perceptrons until predictions are accurate
    for sample := 0; sample < k; sample++ {
        rand.Seed(time.Now().UnixNano())
        newVal := rand.NormFloat64();
        newRBF.weights = append(newRBF.weights, newVal);
    }
    rand.Seed(time.Now().UnixNano())
    newRBF.bias = rand.NormFloat64(); // may consume more memory
    return *newRBF;
}

func rbfFunction(xVal float64, center float64, standardDev float64) float64 {
    denom := 2*(math.Pow(standardDev,float64(2)));
    num := -1*(math.Pow(xVal-center, float64(2)));
    return math.Exp(num/denom);
}

func (rbf *RadialBasis) fitRBF(input []float64, output []float64) {
    rbf.centers, rbf.standardDevs = kMeans(input, rbf.k);
    for epoch := 0; epoch < rbf.epochs; epoch++ {
        for i , val := range input {
            rbfVals := []float64{};
            for j := range rbf.centers {
                rbfVals = append(rbfVals, rbfFunction(val, rbf.centers[j], rbf.standardDevs[j]));
            }
            var result float64;
            for j := range rbfVals {
                result += rbfVals[j]*rbf.weights[j];
            }
            estimate := result + rbf.bias
            //loss := (output[i]-estimate)*(output[i]-estimate);
            //fmt.Println("Loss: ", loss);
            error := (output[i]-estimate); // backwards
            // adjust weights and bias
            for k := range rbf.weights {
                rbf.weights[k] += rbf.learningRate*rbfVals[k]*error;
            }
            rbf.bias += rbf.learningRate*error;
        }
    }
}

func (rbf *RadialBasis) predictRBF(input []float64) []float64 {
    predictions := []float64{};
    for _ , val := range input {
        rbfVals := []float64{};
        for i := range rbf.centers {
            rbfVals = append(rbfVals, rbfFunction(val, rbf.centers[i], rbf.standardDevs[i]));
        }
        var result float64;
        for j := range rbfVals {
            result += rbfVals[j]*rbf.weights[j];
        }
        estimate := result + rbf.bias;
        predictions = append(predictions, estimate);
    }
    return predictions;
}
/* KMeans Clustering - RBF */

func getStandardDev(input []float64, mean float64) float64 {
    var squareDiffs float64;
    for j := range input {
        squareDiffs += math.Pow(input[j]-mean, float64(2));
    }
    return math.Sqrt(squareDiffs/float64(len(input)-1));
}

func getNewCluster(X []float64, clusters []float64, k int) ([]float64, []int, []float64, map[float64]int, map[float64][]float64) {
    clusterIndices := make(map[float64]int);
    for i := range clusters {
        clusterIndices[clusters[i]] = i;
    }
    clustersWithNoNeighbors := []int{};
    standardDevs := make([]float64, len(clusters));
    nearestToClusters := make(map[float64][]float64);
    for i := range X {
        var closerCluster float64;
        minDist := math.MaxFloat64;
        for j := range clusters {
            if (abs(X[i]-clusters[j]) < float64(minDist)) {
                minDist = abs(X[i]-clusters[j]);
                closerCluster = clusters[j]
            }
        }
        nearestToClusters[closerCluster] = append(nearestToClusters[closerCluster], X[i]);
    }
    for cluster, neighbors := range nearestToClusters {
        var mean float64;
        for _, neighbor := range neighbors {
            mean += neighbor;
        }
        if (len(neighbors) > 0) {
            clusters[clusterIndices[cluster]] = mean/float64(len(neighbors));
        }
        if (len(neighbors) < 2) {
            clustersWithNoNeighbors = append(clustersWithNoNeighbors, clusterIndices[cluster]);
        } else {
            standardDevs[clusterIndices[cluster]] = getStandardDev(neighbors, mean/float64(len(neighbors)));
        }
    }
    return clusters, clustersWithNoNeighbors, standardDevs, clusterIndices, nearestToClusters;
}
func linearDifference(l1 []float64, l2 []float64) float64 {
    if (len(l1) != len(l2)) {
        return float64(0);
    } else {
        var diff float64;
        for i := range l1 {
            diff += abs(l1[i]-l2[i]);
        }
        return diff;
    }
}

func kMeans(X []float64, k int) ([]float64, []float64) {
    /* compute standard deviation
       ndarray - M by 1 array inputs
       k - number of clusters
       k by 1 array of cluster centers
    */
    seen := make(map[float64]bool);
    seenIndices := make(map[int]bool);
    var cluster []float64;
    for (len(cluster) < k) {
        rand.Seed(time.Now().UnixNano())
        newIndex := rand.Intn(len(X));
        for (seenIndices[newIndex] || seen[X[newIndex]]) {
            newIndex = rand.Intn(len(X));
        }
        cluster = append(cluster, X[newIndex]);
        seenIndices[newIndex] = true;
        seen[X[newIndex]] = true;
    } // select initial cluster centroids of k clusters
    seen = nil;
    seenIndices = nil;

    var originClusters []float64;
    originClusters = append(originClusters, cluster...);
    standardDevs := make([]float64, k);
    var converged bool = false;
    for (!converged) {
        cluster, _, _, _, _= getNewCluster(X, cluster, k);
        // time complexity O(n) to copy over but required to
        // observe convergence of data
        converged = linearDifference(cluster, originClusters) < math.Pow(float64(10),float64(-20));
        originClusters = []float64{};
        originClusters = append(originClusters, cluster...);
    }
    _, noNeighbors, standardDevs, clusterIndices, nearestToClusters := getNewCluster(X, cluster, k);
    if (len(noNeighbors) != 0) {
        for cluster, neighbors := range nearestToClusters {
            var mean float64 = 0;
            for _, neighbor := range neighbors {
                mean += neighbor;
            }
            if (len(neighbors) >= 2) {
                standardDevs[clusterIndices[cluster]] = getStandardDev(neighbors, mean/float64(len(neighbors)));
            }
        }
    }
    return cluster, standardDevs;
}


// Perceptron implementation
// using an activation function
// signum function best used with linearly separable problems

type PerceptNetwork struct {
    input [][]float64
    weightsLayer1 [][]float64
    weightsLayer2 [][]float64
    firstLayer [][]float64
    output [][] float64
    predictedOutput [][] float64
}

func NewNetwork(input [][]float64, output [][] float64) PerceptNetwork {
    nNetwork := new(PerceptNetwork);
    nNetwork.input = input;
    nNetwork.output = output;
    nNetwork.weightsLayer1 = [][]float64{};
    for r := 0; r < len(input[0]); r++ {
        row := []float64{};
        for c := 0; c < 4; c++ {
            rand.Seed(time.Now().UnixNano());
            newVal := rand.NormFloat64();
            row = append(row, newVal);
        }
        nNetwork.weightsLayer1 = append(nNetwork.weightsLayer1, row);
    }
    nNetwork.weightsLayer2 = [][]float64{};
    for r := 0; r < 4; r++ {
        row := []float64{};
        for c := 0; c < 1; c++ {
            rand.Seed(time.Now().UnixNano());
            val := rand.NormFloat64();
            row = append(row, val);
        }
        nNetwork.weightsLayer2 = append(nNetwork.weightsLayer2, row);
    }
    nNetwork.predictedOutput = [][]float64{};
    return *nNetwork;
} // iteratively undergo some number of epochs

func dotProduct(v1 []float64, v2 []float64) float64 {
    var product float64;
    for i := range v1 {
        product += v1[i]*v2[i];
    }
    return product;
}

func multiply(m1 [][]float64, m2 [][]float64) [][]float64{
    finalRows, finalCols := len(m1), len(m2[0]);
    resultMatrice := [][]float64{};
    for r := 0; r < finalRows; r++ { // initialize matrix
        row := []float64{};
        for c := 0; c < finalCols; c++ {
            m2Col := []float64{};
            for r := 0; r < len(m2); r++ {
                m2Col = append(m2Col, m2[r][c]);
            }
            //fmt.Println(m1,m2Col);
            row = append(row, dotProduct(m1[r], m2Col));
        }
        resultMatrice = append(resultMatrice, row);
    }
    return resultMatrice;
}

func transpose(m [][]float64) [][]float64 {
    newMatrix := [][]float64{};
    for r := 0; r < len(m[0]); r++ {
        row := []float64{};
        for c := 0; c < len(m); c++ {
            row = append(row, m[c][r]);
        }
        newMatrix = append(newMatrix, row);
    }
    return newMatrix;
}

func linearAdd(l1 [][]float64, l2 [][]float64, add bool) [][]float64 {
    if (len(l1) != len(l2)) {
        panic("Unequal lengths");
    } else {
        result := [][]float64{};
        for i := range l1 {
            row := []float64{};
            for j := range l1[i] {
                if (add) {
                    row = append(row, l1[i][j] + l2[i][j]);
                } else {
                    row = append(row, l1[i][j] - l2[i][j]);
                }
            }
            result = append(result, row);
        }
        return result;
    }
}

func constMult(l [][]float64, c float64) [][]float64 {
    newMatrix := [][]float64{};
    for i := range l {
        row := []float64{};
        for col := range l[i] {
            row = append(row, l[i][col]*c);
        }
        newMatrix = append(newMatrix, row);
    }
    return newMatrix;
}

func sigmoid(val float64) float64{
    return float64(1)/(float64(1)+math.Exp(-val));
}

func sigmoidDerivative(val float64) float64 {
    return val*(1-val);
}

func sigmoidDerivativeLst(l [][]float64) [][]float64 {
    newMatrix := [][]float64{};
    for j := range l {
        row := []float64{};
        for c := range l[j] {
            row = append(row, sigmoidDerivative(l[j][c]));
        }
        newMatrix = append(newMatrix, row);
    }
    return newMatrix;
}

func sigmoidLst(l [][]float64) [][]float64 {
    newMatrix := [][]float64{};
    for i := range l {
        row := []float64{};
        for j := range l[i] {
            row = append(row, sigmoid(l[i][j]));
        }
        newMatrix = append(newMatrix, row);
    }
    return newMatrix;
}

func (pnetwork *PerceptNetwork) feedforward() {
    pnetwork.firstLayer = sigmoidLst(multiply(pnetwork.input, pnetwork.weightsLayer1));
    // input and weights w1x1+w2x2...
    pnetwork.predictedOutput = sigmoidLst(multiply(pnetwork.firstLayer, pnetwork.weightsLayer2));
}

func (pnetwork *PerceptNetwork) backpropagate() {
    lossLayer1 := multiply(transpose(pnetwork.input), multiply(multiply(multiply(constMult(linearAdd(pnetwork.output,pnetwork.predictedOutput, false), float64(2)), sigmoidDerivativeLst(pnetwork.predictedOutput)), transpose(pnetwork.weightsLayer2)), sigmoidDerivativeLst(pnetwork.firstLayer))); // derivate output vs predicted
    lossLayer2 := multiply(transpose(pnetwork.firstLayer), multiply(constMult(linearAdd(pnetwork.output,pnetwork.predictedOutput, false), float64(2)), sigmoidDerivativeLst(pnetwork.predictedOutput)));
    pnetwork.weightsLayer1 = linearAdd(pnetwork.weightsLayer1, lossLayer1, true);
    pnetwork.weightsLayer2 = linearAdd(pnetwork.weightsLayer2, lossLayer2, true);
}

func (pnetwork *PerceptNetwork) fitPerceptNetwork(epochs int) {
    for epoch := 0; epoch < epochs; epoch++ {
        pnetwork.feedforward();
        pnetwork.backpropagate();
        if (epoch > 19950) {
            fmt.Println(pnetwork.weightsLayer1,pnetwork.weightsLayer2);
        }
    }
}

func main() {
    in := [][]float64{{0,0,1}, {0,1,1}, {1,0,1}, {1,1,1},{0,1,0}};
    out := [][]float64{{0},{1},{1},{0},{1}};
    newNetwork := NewNetwork(in, out);
    newNetwork.fitPerceptNetwork(20000);
    fmt.Println(newNetwork.predictedOutput);
    // input := []float64{};
    // output := []float64{};
    // var newVal float64;
    // for i := 0; i < 100; i++ {
    //     rand.Seed(time.Now().UnixNano())
    //     newVal = rand.NormFloat64();
    //     input = append(input, newVal);
    //     //noise := rand.NormFloat64();
    //     output = append(output, math.Sin(2*math.Pi*newVal));
    // }
    // myRBF := Constructor(math.Pow(float64(10),float64(-10)), 6, 200);
    // myRBF.fitRBF(input,output);
    // predicted := myRBF.predictRBF(input);
    // fmt.Println(output);
    // fmt.Println(predicted);

}
