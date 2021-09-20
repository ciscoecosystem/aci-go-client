package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (sm *ServiceManager) CreateAnnotation(annotationObject string, fvAnnotationattr models.AnnotationAttributes) (*models.Annotation, error) {
	tagAnnotation := models.NewAnnotation(annotationObject, fvAnnotationattr)
	err := sm.Save(tagAnnotation)
	return tagAnnotation, err
}

func (sm *ServiceManager) UpdateAnnotation(annotationObject string, fvAnnotationattr models.AnnotationAttributes) (*models.Annotation, error) {
	tagAnnotation := models.NewAnnotation(annotationObject, fvAnnotationattr)
  tagAnnotation.Status = "modified"
  err := sm.Save(tagAnnotation)
	return tagAnnotation, err
}

func (sm *ServiceManager) ReadAnnotation(objectName string, key string) (*models.Annotation, error) {
  dn := fmt.Sprintf("%s/annotationKey-[%s]", objectName, key)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	tagAnnotation := models.AnnotationFromContainer(cont)
	return tagAnnotation, nil
}

func (sm *ServiceManager) DeleteAnnotation(objectName string, key string) error {
  dn := fmt.Sprintf("%s/annotationKey-[%s]", objectName, key)
	return sm.DeleteByDn(dn, models.TagAnnotationClassName)
}

func (sm *ServiceManager) ListAnnotation() ([]*models.Annotation, error) {

  baseurlStr := "/api/node/class"
  dnUrl := fmt.Sprintf("%s/tagAnnotation.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.AnnotationListFromContainer(cont)

	return list, err
}
