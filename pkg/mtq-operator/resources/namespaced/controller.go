package namespaced

import (
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	utils2 "kubevirt.io/managed-tenant-quota/pkg/mtq-operator/resources/utils"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sdkapi "kubevirt.io/controller-lifecycle-operator-sdk/api"
)

const (
	controllerResourceName = utils2.ControllerPodName
	SecretResourceName     = "mtq-lock-server-cert"
)

func createMTQControllerResources(args *FactoryArgs) []client.Object {
	return []client.Object{
		createMTQControllerServiceAccount(),
		createControllerRoleBinding(),
		createControllerRole(),
		createMTQControllerDeployment(args.ControllerImage, args.Verbosity, args.PullPolicy, args.ImagePullSecrets, args.PriorityClassName, args.InfraNodePlacement),
	}
}
func createControllerRoleBinding() *rbacv1.RoleBinding {
	return utils2.ResourceBuilder.CreateRoleBinding(controllerResourceName, controllerResourceName, utils2.ControllerServiceAccountName, "")
}
func createControllerRole() *rbacv1.Role {
	rules := []rbacv1.PolicyRule{
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"configmaps",
			},
			Verbs: []string{
				"get",
				"list",
				"watch",
			},
		},
	}
	return utils2.ResourceBuilder.CreateRole(controllerResourceName, rules)
}

func createMTQControllerServiceAccount() *corev1.ServiceAccount {
	return utils2.ResourceBuilder.CreateServiceAccount(controllerResourceName)
}

func createMTQControllerDeployment(image, verbosity, pullPolicy string, imagePullSecrets []corev1.LocalObjectReference, priorityClassName string, infraNodePlacement *sdkapi.NodePlacement) *appsv1.Deployment {
	defaultMode := corev1.ConfigMapVolumeSourceDefaultMode
	deployment := utils2.CreateDeployment(controllerResourceName, utils2.MTQLabel, controllerResourceName, controllerResourceName, imagePullSecrets, 1, infraNodePlacement)
	if priorityClassName != "" {
		deployment.Spec.Template.Spec.PriorityClassName = priorityClassName
	}
	container := utils2.CreateContainer(controllerResourceName, image, verbosity, pullPolicy)
	container.Env = []corev1.EnvVar{
		{
			Name: utils2.InstallerPartOfLabel,
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					APIVersion: "v1",
					FieldPath:  fmt.Sprintf("metadata.labels['%s']", utils2.AppKubernetesPartOfLabel),
				},
			},
		},
		{
			Name: utils2.InstallerVersionLabel,
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					APIVersion: "v1",
					FieldPath:  fmt.Sprintf("metadata.labels['%s']", utils2.AppKubernetesVersionLabel),
				},
			},
		},
	}
	container.ReadinessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			Exec: &corev1.ExecAction{
				Command: []string{"cat", "/tmp/ready"},
			},
		},
		InitialDelaySeconds: 2,
		PeriodSeconds:       5,
		FailureThreshold:    3,
		SuccessThreshold:    1,
		TimeoutSeconds:      1,
	}
	container.Resources = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("50m"),
			corev1.ResourceMemory: resource.MustParse("150Mi"),
		},
	}
	deployment.Spec.Template.Spec.Containers = []corev1.Container{container}
	deployment.Spec.Template.Spec.Volumes = []corev1.Volume{
		{
			Name: "server-cert",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: SecretResourceName,
					Items: []corev1.KeyToPath{
						{
							Key:  "tls.crt",
							Path: "tls.crt",
						},
						{
							Key:  "tls.key",
							Path: "tls.key",
						},
					},
					DefaultMode: &defaultMode,
				},
			},
		},
	}
	return deployment
}
