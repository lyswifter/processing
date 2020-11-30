package model

// import (
// 	"fmt"
// 	"io"

// 	cbg "github.com/whyrusleeping/cbor-gen"
// 	xerrors "golang.org/x/xerrors"
// )

// var _ = xerrors.Errorf

// func (t *UploadStorageClusterResult) MarshalCBOR(w io.Writer) error {
// 	if t == nil {
// 		_, err := w.Write(cbg.CborNull)
// 		return err
// 	}
// 	if _, err := w.Write([]byte{162}); err != nil {
// 		return err
// 	}

// 	scratch := make([]byte, 9)

// 	// t.IsSyncStoreCluster (bool) (bool)
// 	if len("IsSyncStoreCluster") > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field \"IsSyncStoreCluster\" was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("IsSyncStoreCluster"))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, "IsSyncStoreCluster"); err != nil {
// 		return err
// 	}

// 	if err := cbg.WriteBool(w, t.IsSyncStoreCluster); err != nil {
// 		return err
// 	}

// 	// t.State (upload_cluster.TaskState) (string)
// 	if len("State") > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field \"State\" was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("State"))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, "State"); err != nil {
// 		return err
// 	}

// 	if len(t.State) > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field t.State was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.State))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, string(t.State)); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (t *UploadStorageClusterResult) UnmarshalCBOR(r io.Reader) error {
// 	*t = UploadStorageClusterResult{}

// 	br := cbg.GetPeeker(r)
// 	scratch := make([]byte, 8)

// 	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
// 	if err != nil {
// 		return err
// 	}
// 	if maj != cbg.MajMap {
// 		return fmt.Errorf("cbor input should be of type map")
// 	}

// 	if extra > cbg.MaxLength {
// 		return fmt.Errorf("UploadStorageClusterResult: map struct too large (%d)", extra)
// 	}

// 	var name string
// 	n := extra

// 	for i := uint64(0); i < n; i++ {

// 		{
// 			sval, err := cbg.ReadStringBuf(br, scratch)
// 			if err != nil {
// 				return err
// 			}

// 			name = string(sval)
// 		}

// 		switch name {
// 		// t.IsSyncStoreCluster (bool) (bool)
// 		case "IsSyncStoreCluster":

// 			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
// 			if err != nil {
// 				return err
// 			}
// 			if maj != cbg.MajOther {
// 				return fmt.Errorf("booleans must be major type 7")
// 			}
// 			switch extra {
// 			case 20:
// 				t.IsSyncStoreCluster = false
// 			case 21:
// 				t.IsSyncStoreCluster = true
// 			default:
// 				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
// 			}
// 			// t.State (upload_cluster.TaskState) (string)
// 		case "State":

// 			{
// 				sval, err := cbg.ReadStringBuf(br, scratch)
// 				if err != nil {
// 					return err
// 				}

// 				t.State = upload_cluster.TaskState(sval)
// 			}

// 		default:
// 			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
// 		}
// 	}

// 	return nil
// }

// func (t *SectorInfoExtern) MarshalCBOR(w io.Writer) error {
// 	if t == nil {
// 		_, err := w.Write(cbg.CborNull)
// 		return err
// 	}
// 	if _, err := w.Write([]byte{164}); err != nil {
// 		return err
// 	}

// 	scratch := make([]byte, 9)

// 	// t.VdeProgress (int64) (int64)
// 	if len("VdeProgress") > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field \"VdeProgress\" was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("VdeProgress"))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, "VdeProgress"); err != nil {
// 		return err
// 	}

// 	if t.VdeProgress >= 0 {
// 		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.VdeProgress)); err != nil {
// 			return err
// 		}
// 	} else {
// 		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.VdeProgress-1)); err != nil {
// 			return err
// 		}
// 	}

// 	// t.UploadResult (sealing.UploadStorageClusterResult) (struct)
// 	if len("UploadResult") > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field \"UploadResult\" was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("UploadResult"))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, "UploadResult"); err != nil {
// 		return err
// 	}

// 	if err := t.UploadResult.MarshalCBOR(w); err != nil {
// 		return err
// 	}

// 	// t.UpdateTimestamp (int64) (int64)
// 	if len("UpdateTimestamp") > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field \"UpdateTimestamp\" was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("UpdateTimestamp"))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, "UpdateTimestamp"); err != nil {
// 		return err
// 	}

// 	if t.UpdateTimestamp >= 0 {
// 		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.UpdateTimestamp)); err != nil {
// 			return err
// 		}
// 	} else {
// 		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.UpdateTimestamp-1)); err != nil {
// 			return err
// 		}
// 	}

// 	// t.SlaveState (sealing.SectorState) (string)
// 	if len("SlaveState") > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field \"SlaveState\" was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SlaveState"))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, "SlaveState"); err != nil {
// 		return err
// 	}

// 	if len(t.SlaveState) > cbg.MaxLength {
// 		return xerrors.Errorf("Value in field t.SlaveState was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.SlaveState))); err != nil {
// 		return err
// 	}
// 	if _, err := io.WriteString(w, string(t.SlaveState)); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (t *SectorInfoExtern) UnmarshalCBOR(r io.Reader) error {
// 	*t = SectorInfoExtern{}

// 	br := cbg.GetPeeker(r)
// 	scratch := make([]byte, 8)

// 	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
// 	if err != nil {
// 		return err
// 	}
// 	if maj != cbg.MajMap {
// 		return fmt.Errorf("cbor input should be of type map")
// 	}

// 	if extra > cbg.MaxLength {
// 		return fmt.Errorf("SectorInfoExtern: map struct too large (%d)", extra)
// 	}

// 	var name string
// 	n := extra

// 	for i := uint64(0); i < n; i++ {

// 		{
// 			sval, err := cbg.ReadStringBuf(br, scratch)
// 			if err != nil {
// 				return err
// 			}

// 			name = string(sval)
// 		}

// 		switch name {
// 		// t.VdeProgress (int64) (int64)
// 		case "VdeProgress":
// 			{
// 				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
// 				var extraI int64
// 				if err != nil {
// 					return err
// 				}
// 				switch maj {
// 				case cbg.MajUnsignedInt:
// 					extraI = int64(extra)
// 					if extraI < 0 {
// 						return fmt.Errorf("int64 positive overflow")
// 					}
// 				case cbg.MajNegativeInt:
// 					extraI = int64(extra)
// 					if extraI < 0 {
// 						return fmt.Errorf("int64 negative oveflow")
// 					}
// 					extraI = -1 - extraI
// 				default:
// 					return fmt.Errorf("wrong type for int64 field: %d", maj)
// 				}

// 				t.VdeProgress = int64(extraI)
// 			}
// 			// t.UploadResult (sealing.UploadStorageClusterResult) (struct)
// 		case "UploadResult":

// 			{

// 				if err := t.UploadResult.UnmarshalCBOR(br); err != nil {
// 					return xerrors.Errorf("unmarshaling t.UploadResult: %w", err)
// 				}

// 			}
// 			// t.UpdateTimestamp (int64) (int64)
// 		case "UpdateTimestamp":
// 			{
// 				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
// 				var extraI int64
// 				if err != nil {
// 					return err
// 				}
// 				switch maj {
// 				case cbg.MajUnsignedInt:
// 					extraI = int64(extra)
// 					if extraI < 0 {
// 						return fmt.Errorf("int64 positive overflow")
// 					}
// 				case cbg.MajNegativeInt:
// 					extraI = int64(extra)
// 					if extraI < 0 {
// 						return fmt.Errorf("int64 negative oveflow")
// 					}
// 					extraI = -1 - extraI
// 				default:
// 					return fmt.Errorf("wrong type for int64 field: %d", maj)
// 				}

// 				t.UpdateTimestamp = int64(extraI)
// 			}
// 			// t.SlaveState (sealing.SectorState) (string)
// 		case "SlaveState":

// 			{
// 				sval, err := cbg.ReadStringBuf(br, scratch)
// 				if err != nil {
// 					return err
// 				}

// 				t.SlaveState = SectorState(sval)
// 			}

// 		default:
// 			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
// 		}
// 	}

// 	return nil
// }
