/* func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

}

r -->	your controller
ctx -->	request context
req -->	which resource triggered reconcile


// Get Resource
// err := r.Get(ctx, req.NamespacedName, &app)

// Error Handle 


if err != nil {
    return ctrl.Result{}, err
}

** Create / Update Resource
  err = r.Create(ctx, &deployment)
** Update
  err = r.Update(ctx, &deployment)

  return ctrl.Result{RequeueAfter: time.Second * 10}, nil
*/