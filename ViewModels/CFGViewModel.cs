using System.Collections.ObjectModel;
using System.Collections.Generic;

namespace Blob_Editor.ViewModels
{
    public class CFGViewModel : ViewModelBase
    {
        public ObservableCollection<Element> Elements { get; }

        public CFGViewModel(List<Element> elements)
        {
            Elements = new ObservableCollection<Element>(elements);
        }
    }
}