package main

import (
	"flag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {

	var namespace string
	flag.StringVar(&namespace, "namespace", "default", "")
	flag.Parse()

	scheme := runtime.NewScheme()
	restConfig, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))

	// 1. init Manager
	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		Scheme:    scheme,
		Namespace: namespace,
	})
	if err != nil {
		klog.Error("CleanUpController Manager init error.", err)
		os.Exit(1)
	}

	// 2. init Reconciler（Controller）
	reconciler := NewCleanUpController(mgr.GetClient(), mgr.GetScheme())

	err = reconciler.SetupWithManager(mgr)
	if err != nil {
		klog.Error("CleanUpController Reconciler init error.", err)
		os.Exit(1)
	}

	// 3. start Manager
	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		klog.Error("CleanUpController Manager start error.", err)
		os.Exit(1)
	}

}
