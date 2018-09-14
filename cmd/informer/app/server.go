/*
modification history
--------------------
2018/9/7, by lovejoy, create
*/

package app
import (
	"k8s.io/client-go/tools/clientcmd"
	kubeinformers "k8s.io/client-go/informers"
	kubeclientset "k8s.io/client-go/kubernetes"
	restclientset "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"os"
	"time"
	"github.com/golang/glog"
)
func Run() {
	stopCh := make(chan struct{})
	kcfg, err := clientcmd.BuildConfigFromFlags("",os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err)
	}
	kubeClientSet, err2 := kubeclientset.NewForConfig(restclientset.AddUserAgent(kcfg, ""))
	if err2 != nil {
		panic(err2)
	}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClientSet, 30*time.Second)
	// Create pod informer.
	podInformer := kubeInformerFactory.Core().V1().Pods()

	// Set up an event handler for when pod resources change
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    addPod,
		UpdateFunc: updatePod,
		DeleteFunc: deletePod,
	})
	done := make(chan bool, 1)
	go kubeInformerFactory.Start(stopCh)
	<- done


}

func updatePod(old, cur interface{}) {
	glog.Info("updatePod")
	glog.Info(old, cur)
}
func addPod(obj interface{}) {
	glog.Info("addPod")
	glog.Info(obj)
}
func deletePod(obj interface{}) {
	glog.Info("deletePod")
	glog.Info(obj)
}

