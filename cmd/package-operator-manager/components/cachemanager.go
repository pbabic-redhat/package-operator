package components

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"package-operator.run/internal/constants"

	"pkg.package-operator.run/boxcutter/managedcache"
)

func ProvideAccessManager(
	manager ctrl.Manager,
	log logr.Logger,
	restConfig *rest.Config,
	scheme *runtime.Scheme,
) managedcache.ObjectBoundAccessManager[client.Object] {
	mapper := func(
		_ context.Context, _ client.Object,
		c *rest.Config, o cache.Options,
	) (*rest.Config, cache.Options, error) {
		return c, o, nil
	}

	return managedcache.NewObjectBoundAccessManager(
		log,
		mapper,
		restConfig,
		cache.Options{
			Scheme: scheme,
			Mapper: manager.GetRESTMapper(),
			DefaultLabelSelector: labels.SelectorFromSet(labels.Set{
				constants.DynamicCacheLabel: "True",
			}),
		},
	)
}
