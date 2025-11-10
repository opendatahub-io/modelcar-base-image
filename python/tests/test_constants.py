"""Simple tests for modelcar_base_image constants."""
from modelcar_base_image import constants


def test_odh_modelcar_base_image_constant():
    """Test that ODH_MODELCAR_BASE_IMAGE constant is defined."""
    assert constants.ODH_MODELCAR_BASE_IMAGE is not None
    assert isinstance(constants.ODH_MODELCAR_BASE_IMAGE, str)
    assert len(constants.ODH_MODELCAR_BASE_IMAGE) > 0
    assert "odh-modelcar-base-image" in constants.ODH_MODELCAR_BASE_IMAGE


def test_embedded_oci_layout_dir_constant():
    """Test that EMBEDDED_OCI_LAYOUT_DIR constant is defined."""
    assert constants.EMBEDDED_OCI_LAYOUT_DIR is not None
    assert isinstance(constants.EMBEDDED_OCI_LAYOUT_DIR, str)
    assert constants.EMBEDDED_OCI_LAYOUT_DIR == "embedded_oci_layout"

