package lib

// Close - close connection
func (qc *QkvClient) Close() error {
	if err := qc.conn.Close(); err != nil {
		return err
	}
	return nil
}
